package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/bff/initializers"
	"github.com/kuzmindeniss/prost/internal/bff/jwt"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Нет токена авторизации"})
		return
	}

	userId, err := jwt.ReadToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Ошибка при чтении токена: %v", err)})
		return
	}

	repo := repository.New(initializers.DbConn)
	user, err := repo.GetUserById(c, userId)
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
