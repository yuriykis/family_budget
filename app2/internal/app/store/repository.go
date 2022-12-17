package store

import "app/internal/app/model"

type IUserRepository interface {
	Create(model.User) (int, error)
	AuthCheck(string, string) (string, error)
}
