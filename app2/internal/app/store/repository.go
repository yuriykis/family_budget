package store

import "app/internal/app/model"

type IUserRepository interface {
	Create(model.User) error
}
