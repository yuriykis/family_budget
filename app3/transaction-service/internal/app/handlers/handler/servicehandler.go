package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type serviceHealthHandler struct{}

func NewServiceHealthHandler() *serviceHealthHandler {
	return &serviceHealthHandler{}
}

func (s *serviceHealthHandler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
