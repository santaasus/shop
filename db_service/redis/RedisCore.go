package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	root "shop"
)

type Config struct {
	DataBase DataBaseConfig `json:"Database"`
}

type DataBaseConfig struct {
	Postgres RedisConfig `json:"Redis"`
}

type RedisConfig struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Password string `json:"password"`
	Version  int    `json:"version"`
}

const (
	USER     = "user:"
	PRODUCTS = "products"
	ORDER    = "order:"
)

var ctx = context.Background()

func connect() (*redis.Client, error) {
	var config Config

	data, err := root.FileByName("config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%d", config.DataBase.Postgres.Host, config.DataBase.Postgres.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   config.DataBase.Postgres.Version,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

func SaveBy(key string, value any) error {
	rdb, err := connect()
	if err != nil {
		return err
	}

	// Redis can persist string or []byte
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, key, jsonValue, time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetValueBy(key string) (string, error) {
	rdb, err := connect()
	if err != nil {
		return "", err
	}

	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func DeleteValueBy(key string) error {
	rdb, err := connect()
	if err != nil {
		return err
	}

	err = rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
