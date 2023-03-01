package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct{}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{}
}

func (s *ServiceHandler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
