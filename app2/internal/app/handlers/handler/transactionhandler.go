package handler

import (
	"app/internal/app/handlers/requests"
	"app/internal/app/model"
	"app/internal/app/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	store store.IStore
}

func NewTransactionHandler(store store.IStore) *TransactionHandler {
	return &TransactionHandler{store: store}
}

func (tc *TransactionHandler) CreateTransaction(c *gin.Context) {
	var req requests.CreateTransactionRequest
	budgetIDStr := c.Param("budget_id")
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t := model.Transaction{
		Title:      req.Title,
		CategoryID: req.CategoryID,
		Amount:     req.Amount,
		Type:       req.Type,
	}
	budgetID, err := strconv.Atoi(budgetIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := tc.store.Transaction().Create(t, budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (tc *TransactionHandler) GetTransactions(c *gin.Context) {
	budgetIDStr := c.Param("budget_id")
	budgetID, err := strconv.Atoi(budgetIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transactions, err := tc.store.Transaction().FindAll(budgetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
