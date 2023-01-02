package sqlstore

import (
	"app/internal/app/model"
	"app/internal/app/response"
)

type BudgetRepository struct {
	store *Store
}

func (r *BudgetRepository) Create(budget model.Budget, userID uint) (int, error) {

	// sql transaction creates a new budget and user_budget records

	tx, err := r.store.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	err = tx.QueryRow(
		"INSERT INTO budgets (name, description, amount) VALUES ($1, $2, $3) RETURNING id",
		budget.Name,
		budget.Description,
		budget.Amount,
	).Scan(&id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(
		"INSERT INTO user_budgets (budget_id, user_id, ownership, readonly) VALUES ($1, $2, $3, $4)",
		id,
		userID,
		true,
		false,
	)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *BudgetRepository) FindAll() ([]response.BudgetResponse, error) {
	budgetRows, err := r.store.db.Query("SELECT id, name, description, amount FROM budgets")
	if err != nil {
		return nil, err
	}
	defer budgetRows.Close()

	budgets := []response.BudgetResponse{}
	for budgetRows.Next() {
		budget := response.BudgetResponse{}
		err := budgetRows.Scan(&budget.ID, &budget.Name, &budget.Description, &budget.Amount)
		if err != nil {
			return nil, err
		}

		ubsRows, err := r.store.db.Query("SELECT * FROM user_budgets WHERE budget_id = $1", budget.ID)
		if err != nil {
			return nil, err
		}
		defer ubsRows.Close()

		for ubsRows.Next() {

			var id int
			var budgetID int
			var userID uint
			var ownership bool
			var readonly bool
			var tmp string
			var tmp2 string
			err := ubsRows.Scan(&id, &budgetID, &userID, &ownership, &readonly, &tmp, &tmp2)
			if err != nil {
				return nil, err
			}

			budget.AmountLeft = budget.Amount
			budget.Ownership = ownership
			budget.TotalTransactions = 0

			// calculate amount left for budget
			transactionRows, err := r.store.Transaction().FindAll(budget.ID)
			if err != nil {
				return nil, err
			}

			for _, t := range transactionRows {
				amount := t.Amount
				if t.Type == model.TransactionTypeExpense {
					amount = -amount
				}
				budget.AmountLeft += amount
				budget.TotalTransactions++
			}
		}
		budgets = append(budgets, budget)
	}

	return budgets, nil
}

func (r *BudgetRepository) Find(id int) (*model.Budget, error) {
	budget := &model.Budget{}
	err := r.store.db.QueryRow(
		"SELECT id, name, description, amount FROM budgets WHERE id = $1",
		id,
	).Scan(&budget.ID, &budget.Name, &budget.Description, &budget.Amount)

	return budget, err
}

func (r *BudgetRepository) Edit(budget model.Budget) error {
	_, err := r.store.db.Exec(
		"UPDATE budgets SET name = $1, description = $2, amount = $3 WHERE id = $4",
		budget.Name,
		budget.Description,
		budget.Amount,
		budget.ID,
	)

	return err
}

func (r *BudgetRepository) Delete(id int) error {
	tx, err := r.store.db.Begin()
	if err != nil {
		return err
	}

	// delete all transactions for budget
	err = r.store.Transaction().DeleteByBudget(id)

	if err != nil {
		tx.Rollback()
		return err
	}

	// delete all user_budgets for budget
	_, err = r.store.db.Exec("DELETE FROM user_budgets WHERE budget_id = $1", id)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = r.store.db.Exec("DELETE FROM budgets WHERE id = $1", id)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()

	return err
}

func (r *BudgetRepository) Share(budgetID int, userID int) error {
	// check if user already has access to budget
	var count int
	err := r.store.db.QueryRow(
		"SELECT COUNT(*) FROM user_budgets WHERE budget_id = $1 AND user_id = $2",
		budgetID,
		userID,
	).Scan(&count)

	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	_, err = r.store.db.Exec(
		"INSERT INTO user_budgets (budget_id, user_id, ownership) VALUES ($1, $2, $3)", budgetID, userID, false)

	return err
}
