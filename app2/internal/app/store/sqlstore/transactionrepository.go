package sqlstore

import "app/internal/app/model"

type TransactionRepository struct {
	store *Store
}

func (r *TransactionRepository) Create(transaction model.Transaction, budgetID int) (int, error) {

	// check if budget exists
	_, err := r.store.Budget().Find(budgetID)
	if err != nil {
		return 0, err
	}

	err = r.store.db.QueryRow(
		"INSERT INTO transactions (title, budget_id, category_id, amount, type) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		transaction.Title,
		budgetID,
		transaction.CategoryID,
		transaction.Amount,
		transaction.Type,
	).Scan(&transaction.ID)

	return transaction.ID, err
}

func (r *TransactionRepository) FindAll(budgetID int) ([]model.Transaction, error) {
	rows, err := r.store.db.Query("SELECT * FROM transactions WHERE budget_id = $1", budgetID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		transaction := model.Transaction{}
		var tmp string
		var tmp2 string
		err := rows.Scan(
			&transaction.ID,
			&transaction.Title,
			&transaction.BudgetID,
			&transaction.CategoryID,
			&transaction.Amount,
			&transaction.Type,
			&tmp,
			&tmp2,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *TransactionRepository) FindByCategory(categoryID int) ([]model.Transaction, error) {
	rows, err := r.store.db.Query("SELECT * FROM transactions WHERE category_id = $1", categoryID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		transaction := model.Transaction{}
		err := rows.Scan(
			&transaction.ID,
			&transaction.Title,
			&transaction.BudgetID,
			&transaction.CategoryID,
			&transaction.Amount,
			&transaction.Type,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *TransactionRepository) Delete(id int) error {
	_, err := r.store.db.Exec("DELETE FROM transactions WHERE id = $1", id)
	return err
}

func (r *TransactionRepository) DeleteByBudget(budgetID int) error {
	_, err := r.store.db.Exec("DELETE FROM transactions WHERE budget_id = $1", budgetID)
	return err
}
