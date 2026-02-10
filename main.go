package main

import (
	"app-note-go/controllers"
	_ "app-note-go/docs"
	"app-note-go/initializer"
	"app-note-go/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectDB()
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", middleware.VerifyToken, controllers.Ping)
	router.POST("/users/register", controllers.CreateUser)
	router.POST("/users/login", controllers.LoginUser)
	router.Run()
}
