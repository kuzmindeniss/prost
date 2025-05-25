package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/bff/controllers/user"
	"github.com/kuzmindeniss/prost/internal/bff/initializers"
	"github.com/kuzmindeniss/prost/internal/bff/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/sign-up", user.SignUp)
	r.POST("/sign-in", user.SignIn)
	r.POST("/auth", middleware.RequireAuth, user.Auth)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
