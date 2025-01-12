package amqpcore

import (
	"fmt"
	base "shop/amqp_core/base"
)

type ConsumerInfo struct {
	Name      string
	FromQueue string
	Routing   string
}

func (c *ConsumerInfo) Consume() (err error, result []byte) {
	eConfig := &ExchangeMiddleware{
		Kind:    "direct",
		Queue:   c.FromQueue,
		Routing: c.Routing,
	}

	ch, err := eConfig.InitMiddleware()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cConfig := &base.ConsumerConfig{
		Name:  c.Name,
		Queue: c.FromQueue,
		Ch:    ch,
	}

	resultChan := make(chan []byte)
	err = cConfig.Consume(resultChan)
	if err != nil {
		fmt.Println(err.Error())
	}

	result = <-resultChan

	return
}
