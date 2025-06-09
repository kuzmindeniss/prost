package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/bff/jwt"
	"github.com/kuzmindeniss/prost/internal/db"
)

func RequireAuth(c *gin.Context) {
	tokenWithBearer := c.Request.Header.Get("Authorization")

	if tokenWithBearer == "" || !strings.HasPrefix(tokenWithBearer, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Нет токена авторизации"})
		return
	}

	tokenString := strings.TrimPrefix(tokenWithBearer, "Bearer ")

	userId, err := jwt.ReadToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Ошибка при чтении токена: %v", err)})
		return
	}

	user, err := db.Repo.GetUserById(c, userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Ошибка при запросе пользователя"})
		return
	}
	if user.Email == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Пользователь с таким токеном не найден"})
		return
	}

	c.Set("user", user)
	c.Set("tokenString", tokenString)
	c.Next()
}
