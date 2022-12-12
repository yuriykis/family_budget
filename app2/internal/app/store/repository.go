package store

import "app/internal/app/model"

type UserRepository interface {
	Create(model.User) error
}
