package amqpcore

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueConfig struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

func (c *QueueConfig) InitQueue(ch *amqp.Channel) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		c.Name,       // name
		c.Durable,    // durable
		c.AutoDelete, // delete when unused
		c.Exclusive,  // exclusive
		c.NoWait,     // no-wait
		c.Args,       // arguments
	)

	return q, err
}
