package response

type BudgetResponse struct {
	ID                string  `json:"id"                 bson:"_id,omitempty"`
	Name              string  `json:"name"               bson:"budget_name,omitempty"`
	Description       string  `json:"description"        bson:"description,omitempty"`
	Amount            float64 `json:"amount"             bson:"amount,omitempty"`
	Ownership         bool    `json:"ownership"          bson:"ownership,omitempty"`
	AmountLeft        float64 `json:"amount_left"        bson:"amount_left,omitempty"`
	TotalTransactions int     `json:"total_transactions" bson:"total_transactions,omitempty"`
}
