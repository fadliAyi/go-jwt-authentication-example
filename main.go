package main

import (
	"github.com/fadliAyi/go-jwt-authentication-example/controllers"
	"github.com/fadliAyi/go-jwt-authentication-example/database"
	"github.com/fadliAyi/go-jwt-authentication-example/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
