package requests

type CreateTransactionRequest struct {
	Title    string `json:"title"       binding:"required"`
	Category string `json:"category_id" binding:"required"`
	Amount   string `json:"amount"      binding:"required"`
	Type     string `json:"type"        binding:"required"`
}
