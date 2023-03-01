package requests

type CreateTransactionRequest struct {
	Title    string `json:"title" binding:"required"`
	Category int    `json:"category" binding:"required"`
	Amount   string `json:"amount" binding:"required"`
	Type     string `json:"type" binding:"required"`
}
