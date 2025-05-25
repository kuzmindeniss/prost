package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

func Auth(c *gin.Context) {
	user := c.MustGet("user").(repository.User)
	tokenString := c.MustGet("tokenString").(string)

	c.JSON(http.StatusOK, gin.H{
		"user": UserResponse{
			Id:      user.ID,
			Name:    user.Name,
			Surname: user.Surname,
			Email:   user.Email,
			Role:    string(user.Role.UserRoles),
		},
		"token": tokenString,
	})
}
