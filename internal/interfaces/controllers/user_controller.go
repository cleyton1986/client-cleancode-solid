package controllers

import (
	"net/http"

	"github.com/cleyton1986/client-cleancode-solid/internal/entities"
	"github.com/cleyton1986/client-cleancode-solid/internal/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    useCase *usecases.UserUseCase
}

func NewUserController(useCase *usecases.UserUseCase) *UserController {
    return &UserController{
        useCase: useCase,
    }
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    var user entities.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := ctrl.useCase.CreateUser(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) ListUsers(c *gin.Context) {
    users, err := ctrl.useCase.ListUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
    var user entities.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := ctrl.useCase.UpdateUser(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
    var id uint
    if err := c.BindQuery(&id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := ctrl.useCase.DeleteUser(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "success"})
}
