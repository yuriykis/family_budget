package handlers

import "github.com/gin-gonic/gin"

type IServiceHandler interface {
	CheckHealth(c *gin.Context)
}

type ICategoryHandler interface {
	CreateCategory(c *gin.Context)
	FindAllCategories(c *gin.Context)
}
