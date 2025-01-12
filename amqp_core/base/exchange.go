package amqpcore

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type ExchangeConfig struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

func (e *ExchangeConfig) InitExchange(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		e.Name,       // name
		e.Kind,       // type
		e.Durable,    // durable
		e.AutoDelete, // auto-deleted
		e.Internal,   // internal
		e.NoWait,     // no-wait
		e.Args,       // arguments
	)

	if err != nil {
		return err
	}

	return nil
}
