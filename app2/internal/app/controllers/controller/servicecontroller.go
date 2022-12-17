package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceController struct{}

func NewServiceController() *ServiceController {
	return &ServiceController{}
}

func (s *ServiceController) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
