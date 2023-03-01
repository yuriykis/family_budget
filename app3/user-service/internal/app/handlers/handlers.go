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
