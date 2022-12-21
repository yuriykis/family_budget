package controllers

import "github.com/gin-gonic/gin"

type IUserController interface {
	RegisterUser(c *gin.Context)
	AthenticateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserByID(c *gin.Context)
}

type IServiceController interface {
	CheckHealth(c *gin.Context)
}

type IBudgetController interface {
	CreateBudget(c *gin.Context)
	GetAllBudgets(c *gin.Context)
	GetBudgetByID(c *gin.Context)
}
