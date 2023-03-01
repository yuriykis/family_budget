package handlers

import "github.com/gin-gonic/gin"

type IUserHandler interface {
	RegisterUser(c *gin.Context)
	AthenticateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserByID(c *gin.Context)
}

type IServiceHandler interface {
	CheckHealth(c *gin.Context)
}

type IBudgetHandler interface {
	CreateBudget(c *gin.Context)
	GetAllBudgets(c *gin.Context)
	GetBudgetByID(c *gin.Context)
	EditBudget(c *gin.Context)
	DeleteBudget(c *gin.Context)
	ShareBudget(c *gin.Context)
}

type ICategoryHandler interface {
	CreateCategory(c *gin.Context)
	FindAllCategories(c *gin.Context)
}

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
	GetTransactions(c *gin.Context)
}
