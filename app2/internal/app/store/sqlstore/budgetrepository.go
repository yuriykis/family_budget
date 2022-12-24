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

func (r *BudgetRepository) FindAll() ([]model.Budget, error) {
	rows, err := r.store.db.Query("SELECT id, name, description, amount FROM budgets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	budgets := []model.Budget{}
	for rows.Next() {
		budget := model.Budget{}
		err := rows.Scan(&budget.ID, &budget.Name, &budget.Description, &budget.Amount)
		if err != nil {
			return nil, err
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
	_, err := r.store.db.Exec("DELETE FROM budgets WHERE id = $1", id)

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

	_, err = r.store.db.Exec(
		"INSERT INTO user_budgets (budget_id, user_id, ownership) VALUES ($1, $2, $3)", budgetID, userID, false)

	return err
}
