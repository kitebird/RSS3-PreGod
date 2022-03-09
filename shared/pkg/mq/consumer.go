package mq

import (
	"context"
	"errors"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/defers"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

var C Consumer

// Setups a new consumer globally.
func SetupConsumer() error {
	P := &Consumer{
		subscribers: make(map[TopicName]func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)),
	}

	if err := P.Start(); err != nil {
		return err
	}

	return nil
}

type Consumer struct {
	rocketmq.PushConsumer

	subscribers map[TopicName]func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)
}

func (c *Consumer) Subscribe(topic TopicName, f func(context.Context, ...*primitive.MessageExt) error) error {
	if _, ok := c.subscribers[topic]; ok {
		return errors.New("topic already subscribed")
	}

	fn := func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			err := f(ctx, msg)
			if err != nil {
				return consumer.ConsumeRetryLater, err
			}
		}

		return consumer.ConsumeSuccess, nil
	}

	c.subscribers[topic] = fn

	return nil
}

func (c *Consumer) Start() error {
	pc, err := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{config.Config.MQ.NsAddr})),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithNamespace(Namespace),
		consumer.WithGroupName(GroupName),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
	)
	if err != nil {
		return err
	}

	for topic, fn := range c.subscribers {
		if err := c.PushConsumer.Subscribe(string(topic), consumer.MessageSelector{}, fn); err != nil {
			return err
		}
	}

	pc.Start()

	c.PushConsumer = pc

	defers.Register(c.Close)

	return nil
}

// Closes the consumer.
func (c *Consumer) Close() error {
	if err := c.Shutdown(); err != nil {
		return err
	}

	return nil
}
