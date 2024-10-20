package service

import (
	"errors"
	"fmt"
	"net/smtp"
	"shop/notification_service/inner_layer/domain"
	"shop/notification_service/inner_layer/repository"

	domainErrors "github.com/santaasus/errors-handler"
)

type Service struct {
	Repository repository.IRepository
}

func (s *Service) SendMail(userInfo *domain.NotificationInfo) error {
	config, err := s.Repository.GetConfig()
	if err != nil {
		return &domainErrors.AppError{
			Err:  err,
			Type: domainErrors.ValidationError,
		}
	}

	subject := "Subject: Order status\r\n"
	body := fmt.Sprintf("Your order %d was payed.\r\n", userInfo.OrderId)
	message := []byte(subject + "\r\n" + body)

	err = smtp.SendMail(config.Source, nil, config.MailFrom, []string{userInfo.UserEmail}, message)
	if err != nil {
		return &domainErrors.AppError{
			Err:  errors.New("error sending email"),
			Type: domainErrors.NotFound,
		}
	}

	return nil
}
