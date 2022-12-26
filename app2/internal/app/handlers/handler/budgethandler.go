package handler

import (
	"app/internal/app/handlers/requests"
	"app/internal/app/model"
	"app/internal/app/store"
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

func (bc *BudgetHandler) CreateBudget(c *gin.Context) {
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
	id, err := bc.store.Budget().Create(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (bc *BudgetHandler) GetBudgetByID(c *gin.Context) {
	id := c.Param("budget_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	budget, err := bc.store.Budget().Find(id_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budget)
}

func (bc *BudgetHandler) GetAllBudgets(c *gin.Context) {
	budgets, err := bc.store.Budget().FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budgets)
}

func (bc *BudgetHandler) EditBudget(c *gin.Context) {
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
	err = bc.store.Budget().Edit(b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Budget updated"})
}

func (bc *BudgetHandler) DeleteBudget(c *gin.Context) {
	id := c.Param("budget_id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = bc.store.Budget().Delete(id_int)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Budget deleted"})
}

func (bc *BudgetHandler) ShareBudget(c *gin.Context) {
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
	err = bc.store.Budget().Share(id_int, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Budget shared"})
}
