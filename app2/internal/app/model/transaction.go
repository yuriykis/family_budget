package model

type Transaction struct {
	ID         int
	Title      string
	BudgetID   int
	CategoryID int
	Amount     float64
	Type       string
}

// transaction types

const (
	TransactionTypeIncome  = "INCOME"
	TransactionTypeExpense = "EXPENSE"
)
