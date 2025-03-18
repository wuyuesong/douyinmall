package email

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/wuyuesong/douyinmall/app/email/biz/dal/redis"
	"github.com/wuyuesong/douyinmall/app/email/infra/mq"
	"github.com/wuyuesong/douyinmall/app/email/infra/notify"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/email"
)

func ConsumerInit() {
	// Connect to a server

	err := mq.RmqConsumer.Subscribe("EmailTopic", consumer.MessageSelector{}, messageHandler)
	if err != nil {
		klog.Error(fmt.Sprintf("订阅主题失败: %v", err))
	}

	if err := mq.RmqConsumer.Start(); err != nil {
		panic(fmt.Sprintf("启动消费者失败: %v", err))
	}

	server.RegisterShutdownHook(func() {
		// 关闭RocketMQ消费者
		if err := mq.RmqConsumer.Shutdown(); err != nil {
			klog.Error("RocketMQ消费者关闭失败: %v", err)
		} else {
			klog.Info("RocketMQ消费者已优雅关闭")
		}
	})
}

func messageHandler(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		// 提取OpenTelemetry上下文
		carrier := propagation.HeaderCarrier{}
		for k, v := range msg.GetProperties() {
			carrier.Set(k, v)
		}
		ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

		// 解析消息体
		var emailReq email.EmailReq
		if err := json.Unmarshal(msg.Body, &emailReq); err != nil {
			klog.Errorf("消息解析失败: %v | MsgID: %s", err, msg.MsgId)
			return consumer.ConsumeRetryLater, nil
		}

		// 获取业务唯一标识（示例用OrderID）
		businessID := emailReq.GetOrderId() // 假设GetOrderId()方法存在
		if businessID == "" {
			klog.Errorf("消息缺少业务ID | MsgID: %s", msg.MsgId)
			return consumer.ConsumeSuccess, nil // 根据业务需求决定是否重试
		}

		// 构造Redis键
		redisKey := fmt.Sprintf("email:send:%s", businessID)

		// 使用SET NX EX保证原子性操作
		// 设置初始过期时间5分钟，防止处理失败导致永久阻塞
		acquired, err := redis.RedisClient.SetNX(ctx, redisKey, "processing", 5*time.Minute).Result()
		if err != nil {
			klog.Errorf("Redis操作失败: %v | OrderID: %s", err, businessID)
			return consumer.ConsumeRetryLater, nil
		}

		if !acquired {
			// 双重检查：获取当前状态
			status, err := redis.RedisClient.Get(ctx, redisKey).Result()
			if err != nil {
				klog.Infof("获取锁状态失败，重试 | OrderID: %s", businessID)
				return consumer.ConsumeRetryLater, nil
			}

			if status == "processed" {
				klog.Infof("订单已处理，直接跳过 | OrderID: %s", businessID)
				return consumer.ConsumeSuccess, nil
			} else {
				// 状态为 processing，说明有其他消费者正在处理
				klog.Infof("订单正在处理中，等待重试 | OrderID: %s", businessID)
				return consumer.ConsumeRetryLater, nil
			}
		}

		// 执行业务逻辑
		noopEmail := notify.NewNoopEmail()
		if err := noopEmail.Send(&emailReq); err != nil {
			klog.Errorf("邮件发送失败: %v | OrderID: %s", err, businessID)

			// 处理失败：删除键允许重试（或等待自动过期）
			if err := redis.RedisClient.Del(ctx, redisKey).Err(); err != nil {
				klog.Errorf("Redis删除失败: %v | OrderID: %s", err, businessID)
			}
			return consumer.ConsumeRetryLater, nil
		}

		// 处理成功：延长过期时间至24小时
		if err := redis.RedisClient.Set(ctx, redisKey, "processed", 1*time.Hour).Err(); err != nil {
			klog.Errorf("状态持久化失败: %v | OrderID: %s", err, businessID)
			return consumer.ConsumeRetryLater, nil
		}

		klog.Infof("邮件发送成功 | OrderID: %s", businessID)
	}
	return consumer.ConsumeSuccess, nil
}
