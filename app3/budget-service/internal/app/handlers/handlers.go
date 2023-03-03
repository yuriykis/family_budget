package handlers

import "github.com/gin-gonic/gin"

type IBudgetHandler interface {
	CreateBudget(c *gin.Context)
	GetAllBudgets(c *gin.Context)
	GetBudgetByID(c *gin.Context)
	EditBudget(c *gin.Context)
	DeleteBudget(c *gin.Context)
	ShareBudget(c *gin.Context)
}

type IServiceHealthHandler interface {
	CheckHealth(c *gin.Context)
}
