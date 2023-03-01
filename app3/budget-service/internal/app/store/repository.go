package store

import (
	"budgetservice/internal/app/model"
	"budgetservice/internal/app/response"
)

type IBudgetRepository interface {
	Create(model.Budget, uint) (int, error)
	FindAll() ([]response.BudgetResponse, error)
	Find(int) (*model.Budget, error)
	Edit(model.Budget) error
	Delete(int) error
	Share(int, int) error
}
