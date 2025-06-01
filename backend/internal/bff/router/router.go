package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/bff/controllers/applications"
	"github.com/kuzmindeniss/prost/internal/bff/controllers/units"
	"github.com/kuzmindeniss/prost/internal/bff/controllers/user"
	"github.com/kuzmindeniss/prost/internal/bff/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://your-frontend-domain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/sign-up", user.SignUp)
	r.POST("/sign-in", user.SignIn)
	r.POST("/auth", middleware.RequireAuth, user.Auth)

	r.GET("/applications", applications.GetAll)
	r.PATCH("/applications/change-status", applications.ChangeStatus)
	r.DELETE("/applications/delete", applications.Delete)
	r.GET("/units", units.GetAll)

	return r
}
