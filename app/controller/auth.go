package controller

import (
	"errors"
	"mvcgolang/app/database"
	"mvcgolang/app/helper"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var SECRET_KEY = []byte("raswan_sualdi")

func GenerateToken(userID int) (string, error) {
	// claim := jwt.MapClaims{}
	// claim["user_id"] = userID
	jwt.MapClaims{}["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(Encodedtoken string) (*jwt.Token, error) {
	token, err := jwt.Parse(Encodedtoken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid Token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func AuthMiddleWare(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}
	token, err := ValidateToken(tokenString)
	if err != nil {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userID := int(claim["user_id"].(float64))

	user, err := database.GetUserByID(userID)
	if err != nil {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.Set("currentUser", user)
}
