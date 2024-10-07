package adapter

import "shop/order_service/inner_layer/repository"

type BaseAdapter struct {
	Repository repository.IRepository
}
