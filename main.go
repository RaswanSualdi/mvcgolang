package main

import (
	"mvcgolang/app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", controller.AuthMiddleWare, controller.RegisterUser)
	// api.GET("/users", controller.AuthMiddleWare, controller.GetUser)
	api.POST("/sessions", controller.Login)
	api.POST("/files", controller.AuthMiddleWare, controller.UploadFile)
	api.POST("/letter", controller.PostLetter)
	router.Run()
}
