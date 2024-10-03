package adapter

import "shop/products_service/inner_layer/repository"

type BaseAdapter struct {
	Repository repository.IRepository
}
