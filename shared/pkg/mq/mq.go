package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func InitMQProducer(endpoint string, gid string) rocketmq.Producer {
	p, err := rocketmq.NewProducer(
		//producer.WithNameServer(endPoint),
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{endpoint})),
		producer.WithRetry(2),
		producer.WithGroupName(gid),
	)
	if err != nil {
		panic(err)
	}

	return p // TODO: with entity or with pointer ?
}

func InitMQPushConsumer(endpoint string, gid string) rocketmq.PushConsumer {
	c, err := rocketmq.NewPushConsumer(
		//consumer.WithNameServer(endpoint),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{endpoint})),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(gid),
		//consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		//consumer.WithConsumerModel(consumer.BroadCasting),
		// TODO: Copilot generated this (commented above), but what is this?
	)
	if err != nil {
		panic(err)
	}

	return c // TODO: ...?
}
