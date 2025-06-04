package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/bff/controllers"
	"github.com/kuzmindeniss/prost/internal/bff/helpers"
	"github.com/kuzmindeniss/prost/internal/bff/jwt"
	"github.com/kuzmindeniss/prost/internal/db/repository"
	"golang.org/x/crypto/bcrypt"
)

func checkIfUserExists(c *gin.Context, email string) bool {
	user, err := helpers.Repo.GetUserByEmail(c, email)
	if err != nil {
		return false
	}

	return user.Email != ""
}

func SignUp(c *gin.Context) {
	var reqBody struct {
		Name     string `json:"name" binding:"required"`
		Surname  string `json:"surname" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		controllers.HandleValidations(c, err)
		return
	}

	if checkIfUserExists(c, reqBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Пользователь с таким email уже есть",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Ошибка при хешировании пароля",
		})
		return
	}

	createdUser, err := helpers.Repo.CreateUser(c.Request.Context(), repository.CreateUserParams{
		Name:         reqBody.Name,
		Email:        reqBody.Email,
		Surname:      reqBody.Surname,
		PasswordHash: string(hash),
		Role:         "user",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": error.Error(err),
		})
		return
	}

	token, err := jwt.CreateToken(createdUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при создании токена",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": UserResponse{
			Id:      createdUser.ID,
			Name:    createdUser.Name,
			Surname: createdUser.Surname,
			Email:   createdUser.Email,
			Role:    string(createdUser.Role),
		},
		"token": token,
	})
}
