package handler

import (
	"net/http"
	"strconv"
	"transactionservice/internal/app/handlers/requests"
	"transactionservice/internal/app/model"
	"transactionservice/internal/app/store"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	store store.IStore
}

func NewTransactionHandler(store store.IStore) *TransactionHandler {
	return &TransactionHandler{store: store}
}

func (th *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req requests.CreateTransactionRequest
	budgetIDStr := c.Param("budget_id")
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// convert amount to float
	amountFloat, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := model.Transaction{
		Title:      req.Title,
		CategoryID: req.Category,
		Amount:     amountFloat,
		Type:       req.Type,
	}
	budgetID, err := strconv.Atoi(budgetIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := th.store.Transaction().Create(t, budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (th *TransactionHandler) GetTransactions(c *gin.Context) {
	budgetIDStr := c.Param("budget_id")
	budgetID, err := strconv.Atoi(budgetIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transactions, err := th.store.Transaction().FindAll(budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
