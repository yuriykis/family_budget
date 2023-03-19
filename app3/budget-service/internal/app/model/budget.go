package model

type Budget struct {
	ID          string  `json:"id"          bson:"_id,omitempty"`
	Name        string  `json:"name"        bson:"budget_name,omitempty"`
	Description string  `json:"description" bson:"description,omitempty"`
	Amount      float64 `json:"amount"      bson:"amount,omitempty"`
}

type UserBudget struct {
	UserID    int  `json:"user_id"   bson:"user_id"`
	BudgetID  int  `json:"budget_id" bson:"budget_id"`
	Ownership bool `json:"ownership" bson:"ownership"`
	Readonly  bool `json:"readonly"  bson:"readonly"`
}
