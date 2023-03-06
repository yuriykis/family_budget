package store

import (
	"userservice/internal/app/model"
)

type IUserRepository interface {
	Create(model.User) (string, error)
	Update(model.User) error
	Delete(string) error
	Find(string) (*model.User, error)
	AuthCheck(string, string) (string, error)
	FindAll() ([]model.User, error)
}
