package controller

import (
	"app/internal/app/controllers"
	"app/internal/app/model"
	"app/internal/app/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BudgetController struct {
	store store.IStore
}

func NewBudgetController(store store.IStore) *BudgetController {
	return &BudgetController{store: store}
}

func (bc *BudgetController) CreateBudget(c *gin.Context) {
	var req controllers.CreateBudgetRequest
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

func (bc *BudgetController) GetBudgetByID(c *gin.Context) {
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

func (bc *BudgetController) GetAllBudgets(c *gin.Context) {
	budgets, err := bc.store.Budget().FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budgets)
}
