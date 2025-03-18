package mq

import (
	"context"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/wuyuesong/douyinmall/app/payment/conf"
)

var (
	RmqProducer rocketmq.Producer
	err         error
)

func InitRocketMQ() {
	// 创建生产者配置（默认使用 NameServer 地址）
	RmqProducer, err = rocketmq.NewProducer(
		producer.WithNameServer([]string{conf.GetConf().RocketMQ.NamesrvAddrs[0]}), // NameServer 地址
		producer.WithRetry(2),                        // 发送失败重试次数
		producer.WithGroupName("EmailProducerGroup"), // 生产者组名
	)
	if err != nil {
		panic("Failed to create producer: " + err.Error())
	}

	// 启动生产者
	if err = RmqProducer.Start(); err != nil {
		panic("Failed to start producer: " + err.Error())
	}

	// 发送测试消息验证连接
	testMsg := primitive.NewMessage("ConnectionTestTopic", []byte("Init Test Message"))
	_, err = RmqProducer.SendSync(context.Background(), testMsg)
	if err != nil {
		panic("RocketMQ connection test failed: " + err.Error())
	}
}
