package controller

import (
	"fmt"
	"mvcgolang/app/database"
	"mvcgolang/app/helper"
	"mvcgolang/app/model"
	"mvcgolang/app/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var input request.RegisterUserInput
	err := c.BindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register Account Failed ", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user := model.User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(passwordHash),
		Role:         "user",
	}

	newUser, err := database.SaveUser(user)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register Account Failed ", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register Account Failed ", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := helper.FormatUser(newUser, token)
	responseData := helper.APIResponse("Account Has been Registered ", http.StatusOK, "success", formatter)
	c.JSON(200, responseData)
	// c.MustGet("currentUser")

}

func Login(c *gin.Context) {
	var input request.Login
	err := c.BindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login Gagal ", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	email := input.Email
	password := input.Password

	loggedInUser, err := database.Login(email, password)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse("Register Account Failed ", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := helper.FormatUser(loggedInUser, token)
	response := helper.APIResponse("SuccessFully logged in", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// userID := 3
	currentUser := c.MustGet("currentUser").(model.User)
	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = database.SaveFile(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image gagal", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successFully uploaded", http.StatusOK, "Success", data)
	c.JSON(http.StatusOK, response)
}

func GetUser(c *gin.Context) {
	var users []model.User
	GetUser, err := database.GetUser(users)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed Get Users ", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	responseData := helper.APIResponse("Account Has been Registered ", http.StatusOK, "success", GetUser)
	c.JSON(200, responseData)
}
