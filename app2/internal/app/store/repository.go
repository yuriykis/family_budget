package store

import "app/internal/app/model"

type IUserRepository interface {
	Create(model.User) (int, error)
	Find(int) (*model.User, error)
	AuthCheck(string, string) (string, error)
}

type IBudgetRepository interface {
	Create(model.Budget) (int, error)
}
