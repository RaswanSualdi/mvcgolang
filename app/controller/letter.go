package controller

import (
	"mvcgolang/app/database"
	"mvcgolang/app/helper"
	"mvcgolang/app/model"
	"mvcgolang/app/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostLetter(c *gin.Context) {
	var input request.InputLetter
	err := c.BindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Post Letter failed ", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	id := uuid.New()
	letter := model.Letter{
		Uuid:        id.String(),
		MdLettersId: 1,
		Letter:      "No.001/UPA",
		Description: input.Description,
		Link:        input.Link,
		Slug:        "surat-upsns",
		DateLetter:  "01-01-2000",
		YearLetter:  "2000",
		MonthLetter: "January",
		NoLetter:    "01",
		Images:      "images/photo1.jpg",
	}

	newLetter, err := database.SaveLetter(letter)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Create Letter Failed ", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// formatter := helper.FormatUser(newLetter, token)
	responseData := helper.APIResponse("Letter Has been Created ", http.StatusOK, "success", newLetter)
	c.JSON(200, responseData)

}


func ServiceLetter()
