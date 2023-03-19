package store

import (
	"budgetservice/internal/app/model"
	"budgetservice/internal/app/response"
)

type IBudgetRepository interface {
	Create(model.Budget, string) (string, error)
	FindAll() ([]response.BudgetResponse, error)
	Find(string) (*response.BudgetResponse, error)
	Edit(model.Budget) error
	Delete(int) error
	Share(int, int) error
}
