package main

import (
	"app-note-go/controllers"
	_ "app-note-go/docs"
	"app-note-go/initializer"
	"app-note-go/middleware"
	"app-note-go/migration"

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
	migration.Migrate()

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/users/register", controllers.CreateUser)
	router.POST("/users/login", controllers.LoginUser)

	router.POST("/notes/create", middleware.VerifyToken, controllers.CreateNote)
	router.GET("/notes/pagination", middleware.VerifyToken, controllers.GetNotePagination)
	router.GET("/notes/:id", middleware.VerifyToken, controllers.GetNote)
	router.PUT("/notes/update", middleware.VerifyToken, controllers.UpdateNote)
	router.DELETE("/notes/delete/:id", middleware.VerifyToken, controllers.DeleteNote)

	router.Run()
}
