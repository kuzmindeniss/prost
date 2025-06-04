package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

func RequireAuthAdmin(c *gin.Context) {
	user := c.MustGet("user").(repository.User)

	if user.Role != repository.UserRolesAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "У вас нет доступа"})
		return
	}

	c.Next()
}
