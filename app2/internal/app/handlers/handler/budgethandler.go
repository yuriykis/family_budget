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
	b := model.Budget{
		Name:        req.Name,
		Description: req.Description,
		Amount:      req.Amount,
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
