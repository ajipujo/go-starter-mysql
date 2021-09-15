package handler

import (
	"gostart/auth"
	"gostart/helper"
	"gostart/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	input := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Register Account Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Register Account Failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helper.ApiResponse("Register Account Failed", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatterUser := user.FormatterUser(newUser, token)
	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatterUser)
	c.JSON(http.StatusOK, response)
}
