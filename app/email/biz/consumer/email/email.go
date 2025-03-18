package email

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/wuyuesong/douyinmall/app/email/infra/mq"
	"github.com/wuyuesong/douyinmall/app/email/infra/notify"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/wuyuesong/douyinmall/rpc_gen/kitex_gen/email"
)

func ConsumerInit() {
	// Connect to a server

	err := mq.RmqConsumer.Subscribe("EmailTopic", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				carrier := propagation.HeaderCarrier{}
				for k, v := range msg.GetProperties() {
					carrier.Set(k, v)
				}
				ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

				var emailReq email.EmailReq
				if err := json.Unmarshal(msg.Body, &emailReq); err != nil {
					klog.Error("消息解析失败: %v | MsgID: %s\n", err, msg.MsgId)
					return consumer.ConsumeRetryLater, nil // 触发重试
				}

				noopEmail := notify.NewNoopEmail()
				_ = noopEmail.Send(&emailReq)
			}
			return consumer.ConsumeSuccess, nil
		})
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
