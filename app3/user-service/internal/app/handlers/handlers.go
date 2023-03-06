package handlers

import "github.com/gin-gonic/gin"

type IUserHandler interface {
	RegisterUser(c *gin.Context)
	AthenticateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type IServiceHealthHandler interface {
	CheckHealth(c *gin.Context)
}
