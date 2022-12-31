package store

import (
	"app/internal/app/model"
	"app/internal/app/response"
)

type IUserRepository interface {
	Create(model.User) (int, error)
	Find(int) (*model.User, error)
	AuthCheck(string, string) (string, error)
	FindAll() ([]model.User, error)
}

type IBudgetRepository interface {
	Create(model.Budget, uint) (int, error)
	FindAll() ([]response.BudgetResponse, error)
	Find(int) (*model.Budget, error)
	Edit(model.Budget) error
	Delete(int) error
	Share(int, int) error
}

type ICategoryRepository interface {
	Create(model.Category) (int, error)
	FindAll() ([]model.Category, error)
	Find(int) (*model.Category, error)
}

type ITransactionRepository interface {
	Create(model.Transaction, int) (int, error)
	FindAll(int) ([]model.Transaction, error)
	FindByCategory(int) ([]model.Transaction, error)
}
