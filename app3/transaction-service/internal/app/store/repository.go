package store

import (
	"transactionservice/internal/app/model"
)

type ITransactionRepository interface {
	Create(model.Transaction, int) (int, error)
	Find(int) (*model.Transaction, error)
	FindAll(int) ([]model.Transaction, error)
	FindByCategory(int) ([]model.Transaction, error)
	Delete(int) error
	DeleteByBudget(int) error
}
