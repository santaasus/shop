package adapter

import (
	repository "shop/user_service/inner_layer/repository/user"
)

type BaseAdapter struct {
	Repository repository.IRepository
}
