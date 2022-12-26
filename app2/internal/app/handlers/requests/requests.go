package requests

type RegisterUserRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type AuthRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateBudgetRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateTransactionRequest struct {
	Title      string  `json:"title" binding:"required"`
	CategoryID int     `json:"category_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
	Type       string  `json:"type" binding:"required"`
}

type EditBudgetRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
}

type ShareBudgetRequest struct {
	UserID int `json:"user_id" binding:"required"`
}
