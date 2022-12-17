package controller

import (
	"app/internal/app/controllers"
	"app/internal/app/model"
	"app/internal/app/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	store store.IStore
}

func NewsUserController(store store.IStore) *UserController {
	return &UserController{store: store}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var req controllers.RegisterUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
	id, err := uc.store.User().Create(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (uc *UserController) AthenticateUser(c *gin.Context) {
	var req controllers.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u := model.User{}

	u.Email = req.Email
	u.Password = req.Password

	token, err := uc.store.User().AuthCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
