package mq

import (
	"context"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/defers"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var P Producer

// Setups a new producer globally.
func SetupProducer() error {
	P := &Producer{}

	if err := P.Start(); err != nil {
		return err
	}

	return nil
}

type Producer struct {
	rocketmq.Producer
}

// Starts the producer.
func (p *Producer) Start() error {
	pp, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{config.Config.MQ.NsAddr})),
		producer.WithRetry(2),
		producer.WithNamespace(Namespace),
		producer.WithGroupName(GroupName),
		producer.WithCompressLevel(3),
	)
	if err != nil {
		return err
	}

	if err := pp.Start(); err != nil {
		return err
	}

	p.Producer = pp
	defers.Register(p.Close)

	return nil
}

// Closes the producer.
func (p *Producer) Close() error {
	if err := p.Shutdown(); err != nil {
		return err
	}

	return nil
}

// Sends a message.
func (pc *Producer) Send(topic TopicName, msg []byte) (*primitive.SendResult, error) {
	m := primitive.NewMessage(string(topic), msg)

	res, err := pc.SendSync(context.Background(), m)
	if err != nil {
		return res, err
	}

	return res, nil
}
