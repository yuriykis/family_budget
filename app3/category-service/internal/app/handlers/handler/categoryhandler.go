package handler

import (
	"categoryservice/internal/app/handlers/requests"
	"categoryservice/internal/app/model"
	"categoryservice/internal/app/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	store store.IStore
}

func NewCategoryHandler(store store.IStore) *CategoryHandler {
	return &CategoryHandler{store: store}
}

func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	var req requests.CreateCategoryRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cate := model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	id, err := ch.store.Category().Create(cate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})

}

func (ch *CategoryHandler) FindAllCategories(c *gin.Context) {
	categories, err := ch.store.Category().FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}
