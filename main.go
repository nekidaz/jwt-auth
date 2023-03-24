package main

import (
	"github.com/gin-gonic/gin"
	"jwt-auth/controllers"
	"jwt-auth/initializers"
	"jwt-auth/middleware"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "404 Not Found"})
	})

	// авторизация
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	auth := router.Group("/")
	auth.Use(middleware.RequireAuth)
	{
		auth.GET("/admin", controllers.WelcomePage)
		auth.GET("/validate", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(200, gin.H{
				"message": user,
			})
		})
	}
	router.Run()
}
