package amqpcore

import (
	"context"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ProducerConfig struct {
	Ch       *amqp.Channel
	Exchange string
	Routing  string
	Msg      amqp.Publishing
}

func (p *ProducerConfig) Publish() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := p.Ch.PublishWithContext(
		ctx,
		p.Exchange,
		p.Routing,
		false,
		false,
		p.Msg,
	)

	if err != nil {
		return err
	}

	fmt.Printf("Producer sent %v", p.Msg.Body)

	return nil
}
