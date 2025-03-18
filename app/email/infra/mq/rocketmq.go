package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/wuyuesong/douyinmall/app/email/conf"
)

var RmqConsumer rocketmq.PushConsumer

func InitRocketMQConsumer() {
	var err error

	rlog.SetLogLevel("warn")

	RmqConsumer, err = rocketmq.NewPushConsumer(
		consumer.WithGroupName("EmailConsumerGroup"),
		consumer.WithNameServer([]string{conf.GetConf().RocketMQ.NamesrvAddrs[0]}),
		consumer.WithConsumerModel(consumer.Clustering), // 集群模式
		consumer.WithMaxReconsumeTimes(16),              // 最大重试次数
		consumer.WithConsumeGoroutineNums(5),            // 消费者协程数
	)
	if err != nil {
		klog.Info("Failed to create consumer: %v", err)
	}
}
