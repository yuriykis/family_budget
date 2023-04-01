package model

type Transaction struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	BudgetID   string  `json:"budget_id"`
	CategoryID string  `json:"category_id"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"type"`
}

// transaction types

const (
	TransactionTypeIncome  = "INCOME"
	TransactionTypeExpense = "EXPENSE"
)
