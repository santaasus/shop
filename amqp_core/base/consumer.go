package amqpcore

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumerConfig struct {
	Name  string
	Queue string
	Ch    *amqp.Channel
}

func (c *ConsumerConfig) Consume(chanBody chan []byte) error {
	msgs, err := c.Ch.Consume(
		c.Queue,
		c.Name, // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Consumer received a message: %s", d.Body)

			chanBody <- d.Body
		}
	}()

	<-chanBody

	return nil
}
