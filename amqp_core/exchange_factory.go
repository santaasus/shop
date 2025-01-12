package amqpcore

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	base "shop/amqp_core/base"
)

type ExchangeMiddleware struct {
	Kind    string
	Queue   string
	Routing string
}

func (e *ExchangeMiddleware) InitMiddleware() (*amqp.Channel, error) {
	ch, err := Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	eConfig := &base.ExchangeConfig{
		Name:       BASE_EXCHANGE,
		Kind:       e.Kind,
		Durable:    false,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}

	err = eConfig.InitExchange(ch)
	if err != nil {
		fmt.Println(err.Error())
	}

	qConfig := &base.QueueConfig{
		Name: e.Queue,
	}

	q, err := qConfig.InitQueue(ch)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ch.QueueBind(
		q.Name,        // queue name
		e.Routing,     // routing key
		BASE_EXCHANGE, // exchange
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	return ch, err
}
