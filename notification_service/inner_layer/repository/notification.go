package repository

import (
	"encoding/json"
	"os"
	"shop/notification_service/inner_layer/domain"
	// root "shop"
)

type IRepository interface {
	GetConfig() (*domain.SMTP, error)
}

type Repository struct {
}

func (Repository) GetConfig() (*domain.SMTP, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config *domain.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config.SMTP, nil
}
