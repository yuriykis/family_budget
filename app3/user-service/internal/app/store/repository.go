package store

import (
	"userservice/internal/app/model"
)

type IUserRepository interface {
	Create(model.User) (string, error)
	Update(model.User) error
	Delete(int) error
	Find(int) (*model.User, error)
	AuthCheck(string, string) (string, error)
	FindAll() ([]model.User, error)
}
