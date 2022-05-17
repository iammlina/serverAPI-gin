package main

import (
	"github.com/gin-gonic/gin"
	"server/controllers"
	"server/database"
	"server/middlewares"
)

func main() {
	// Initialize Database
	database.Connect("test:test@tcp(localhost:3306)/testAPI?parseTime=true")
	database.Migrate()

	//Initialize Router
	router := initRouter()
	router.Run(":8000")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/token", controllers.GenerateToken)
		api.GET("/users", controllers.GetUsers)
		api.POST("/user/register", controllers.RegisterUser)
		api.GET("/logout", controllers.Logout).Use(middlewares.Auth())
		//api.POST("/login", controllers.Login)

		// Test
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}

	}
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
