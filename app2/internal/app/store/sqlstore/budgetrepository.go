package sqlstore

import "app/internal/app/model"

type BudgetRepository struct {
	store *Store
}

func (r *BudgetRepository) Create(budget model.Budget) (int, error) {

	err := r.store.db.QueryRow(
		"INSERT INTO budgets (name, description, amount) VALUES ($1, $2, $3) RETURNING id",
		budget.Name,
		budget.Description,
		budget.Amount,
	).Scan(&budget.ID)

	return budget.ID, err
}
