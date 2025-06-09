package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/bff/controllers"
	"github.com/kuzmindeniss/prost/internal/bff/jwt"
	"github.com/kuzmindeniss/prost/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *gin.Context) {
	var reqBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		controllers.HandleValidations(c, err)
		return
	}

	user, err := db.Repo.GetUserByEmail(c, reqBody.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Неверные данные",
		})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(reqBody.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные",
		})
		return
	}

	token, err := jwt.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при создании токена",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": UserResponse{
			Id:      user.ID,
			Name:    user.Name,
			Surname: user.Surname,
			Email:   user.Email,
			Role:    string(user.Role),
		},
		"token": token,
	})
}
