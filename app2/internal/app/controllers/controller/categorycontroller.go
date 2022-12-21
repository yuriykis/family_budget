package controller

import (
	"app/internal/app/controllers"
	"app/internal/app/model"
	"app/internal/app/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	store store.IStore
}

func NewCategoryController(store store.IStore) *CategoryController {
	return &CategoryController{store: store}
}

func (cc *CategoryController) CreateCategory(c *gin.Context) {
	var req controllers.CreateCategoryRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cate := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	id, err := cc.store.Category().Create(cate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})

}
