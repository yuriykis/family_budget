package handler

import (
	"net/http"
	"userservice/internal/app/handlers/requests"
	"userservice/internal/app/model"
	"userservice/internal/app/store"
	"userservice/internal/app/utils/token"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	store store.IStore
}

func NewsUserHandler(store store.IStore) *UserHandler {
	return &UserHandler{store: store}
}

func (uh *UserHandler) RegisterUser(c *gin.Context) {
	var req requests.RegisterUserRequest
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
	id, err := uh.store.User().Create(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (uh *UserHandler) AthenticateUser(c *gin.Context) {
	var req requests.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u := model.User{}

	u.Email = req.Email
	u.Password = req.Password

	token, err := uh.store.User().AuthCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access": token})
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	users, err := uh.store.User().FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("user_id")
	user, err := uh.store.User().Find(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id_param := c.Param("user_id")
	if id != id_param {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req requests.UpdateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := model.User{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
	err = uh.store.User().Update(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("user_id")
	err := uh.store.User().Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
