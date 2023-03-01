package model

type Transaction struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	BudgetID   int     `json:"budget_id"`
	CategoryID int     `json:"category_id"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"type"`
}

// transaction types

const (
	TransactionTypeIncome  = "INCOME"
	TransactionTypeExpense = "EXPENSE"
)
