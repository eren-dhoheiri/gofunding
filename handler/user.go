package handler

import (
	"backend_funding/helper"
	"backend_funding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewsUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

//	tangkap input dari user
//	mapping input dari user ke struct RegisterUserInput
//	struct diatas kita passing sebagai parameter service
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "Failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser, "tokentokentokentokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

//User memasukan input (email & password)
//input ditangkap handler
//mapping dari input user ke input struct
//input struct passing ke service
//di service mencari dg bantuan repository user dengan email x
//mencocokan password
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Your email or password is incorrect", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Your email or password is incorrect", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentokentokentoken")

	response := helper.APIResponse("Login successfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}


// ada input email dari user
// input email di mapping ke struct input
// struct input dipassing ke service
// service akan memanggil repository -email sudah ada atau blm
// repository qeuery ke db
func (h *userHandler) CheckEmailAvailability (c *gin.Context)  {
	var input user.CheckEmailInput 

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Your email has been registered", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Your email has been registered", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available" : IsEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if IsEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
		c.JSON(http.StatusOK, response)

	
}