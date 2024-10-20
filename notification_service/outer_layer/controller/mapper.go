package controller

import "shop/notification_service/inner_layer/domain"

func (n *NotificationRequest) MapToDomain() *domain.NotificationInfo {
	return &domain.NotificationInfo{
		OrderId:   n.OrderId,
		UserEmail: n.UserEmail,
	}
}
