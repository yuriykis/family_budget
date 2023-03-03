package handlers

import "github.com/gin-gonic/gin"

type IServiceHealthHandler interface {
	CheckHealth(c *gin.Context)
}

type ICategoryHandler interface {
	CreateCategory(c *gin.Context)
	FindAllCategories(c *gin.Context)
}
