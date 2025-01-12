package amqpcore

import (
	"encoding/json"
	"fmt"

	root "shop"

	amqp "github.com/rabbitmq/amqp091-go"
)

type config struct {
	Info struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
	} `json:"AMQP"`
}

func Start() (*amqp.Channel, error) {
	file, err := root.FileByName("config.json")
	if err != nil {
		return nil, err
	}

	var config config

	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%d", config.Info.User, config.Info.Password, config.Info.Host, config.Info.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	// defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	// defer ch.Close()

	return ch, nil
}
