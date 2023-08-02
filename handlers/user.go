package handlers

import (
	"net/http"
	"startup-apps/helper"
	"startup-apps/users"

	"github.com/gin-gonic/gin"
)

type userHanlder struct {
	userService users.Service
}

func NewUserHandler(userService users.Service) *userHanlder {
	return &userHanlder{userService}
}

func (h *userHanlder) RegisterUser(c *gin.Context)  {
	// Tangkap input dari user
	// Map input user ke register user input
	// Struct di atas passing ke service parameter
	var input users.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		
		response := helper.ResponseApi("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.ResponseApi("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := users.FormatUser(newUser, "sfdjhjhkjhdKHKKhjkd63784627346923DSKJFHKAJSD")

	response := helper.ResponseApi("Account Has Been Registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}