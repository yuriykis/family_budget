package response

type BudgetResponse struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Amount            float64 `json:"amount"`
	Ownership         bool    `json:"ownership"`
	AmountLeft        float64 `json:"amount_left"`
	TotalTransactions int     `json:"total_transactions"`
}
