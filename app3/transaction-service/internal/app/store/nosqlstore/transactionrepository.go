package nosqlstore

import "transactionservice/internal/app/model"

type TransactionRepository struct {
	store *Store
}

func (r *TransactionRepository) Create(transaction model.Transaction, budgetID int) (int, error) {

	panic("not implemented")
}

func (r *TransactionRepository) FindAll(budgetID int) ([]model.Transaction, error) {
	panic("not implemented")
}

func (r *TransactionRepository) FindByCategory(categoryID int) ([]model.Transaction, error) {
	panic("not implemented")
}

func (r *TransactionRepository) Delete(id int) error {
	panic("not implemented")
}

func (r *TransactionRepository) DeleteByBudget(budgetID int) error {
	panic("not implemented")
}

func (r *TransactionRepository) Find(id int) (*model.Transaction, error) {
	panic("not implemented")
}
