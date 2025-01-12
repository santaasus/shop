package amqpcore

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	base "shop/amqp_core/base"
)

type PublishBody struct {
	ContentType string
	Body        []byte
}

func (b *PublishBody) Publish(routingName string) (err error) {
	ch, err := Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pConfig := &base.ProducerConfig{
		Ch:       ch,
		Exchange: BASE_EXCHANGE,
		Routing:  routingName,
		Msg: amqp.Publishing{
			ContentType: b.ContentType,
			Body:        b.Body,
		},
	}

	err = pConfig.Publish()
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}
