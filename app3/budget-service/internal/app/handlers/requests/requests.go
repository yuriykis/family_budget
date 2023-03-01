package requests

type CreateBudgetRequest struct {
	Name        string `json:"name"        binding:"required"`
	Description string `json:"description" binding:"required"`
	Amount      string `json:"amount"      binding:"required"`
}

type EditBudgetRequest struct {
	Name        string  `json:"name"        binding:"required"`
	Description string  `json:"description" binding:"required"`
	Amount      float64 `json:"amount"      binding:"required"`
}

type ShareBudgetRequest struct {
	UserID int `json:"user_id" binding:"required"`
}
