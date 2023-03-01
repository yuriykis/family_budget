package handler

import (
	"budgetservice/internal/app/handlers/requests"
	"budgetservice/internal/app/model"
	"budgetservice/internal/app/store"
	"budgetservice/internal/app/utils/token"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	store store.IStore
}

func NewBudgetHandler(store store.IStore) *BudgetHandler {
	return &BudgetHandler{store: store}
}

func (bh *BudgetHandler) CreateBudget(c *gin.Context) {
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var req requests.CreateBudgetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	amountFloat, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := model.Budget{
		Name:        req.Name,
		Description: req.Description,
		Amount:      amountFloat,
	}
	id, err := bh.store.Budget().Create(b, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (bh *BudgetHandler) GetBudgetByID(c *gin.Context) {
	id := c.Param("budget_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	budget, err := bh.store.Budget().Find(id_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budget)
}

func (bh *BudgetHandler) GetAllBudgets(c *gin.Context) {
	budgets, err := bh.store.Budget().FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budgets)
}

func (bh *BudgetHandler) EditBudget(c *gin.Context) {
	id := c.Param("budget_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req requests.EditBudgetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := model.Budget{
		ID:          id_int,
		Name:        req.Name,
		Description: req.Description,
		Amount:      req.Amount,
	}
	err = bh.store.Budget().Edit(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Budget updated"})
}

func (bh *BudgetHandler) DeleteBudget(c *gin.Context) {
	id := c.Param("budget_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = bh.store.Budget().Delete(id_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Budget deleted"})
}

func (bh *BudgetHandler) ShareBudget(c *gin.Context) {
	id := c.Param("budget_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req requests.ShareBudgetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = bh.store.Budget().Share(id_int, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Budget shared"})
}
