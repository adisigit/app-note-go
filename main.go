package main

import (
	"app-note-go/controllers"
	_ "app-note-go/docs"
	"app-note-go/initializer"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", controllers.Ping)
	router.POST("/users/register", controllers.CreateUser)
	router.POST("/users/login", controllers.LoginUser)
	router.Run()
}
