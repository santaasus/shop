package adapter

import "shop/notification_service/inner_layer/repository"

type BaseAdapter struct {
	Repository repository.IRepository
}
