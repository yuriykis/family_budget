package handlers

import "github.com/gin-gonic/gin"

type IServiceHealthHandler interface {
	CheckHealth(c *gin.Context)
}

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
	GetTransactions(c *gin.Context)
}
