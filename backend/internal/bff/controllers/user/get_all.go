package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/db"
)

func GetAll(c *gin.Context) {
	users, err := db.Repo.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при получении пользователей: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
