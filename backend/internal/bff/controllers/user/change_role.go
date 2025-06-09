package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kuzmindeniss/prost/internal/bff/controllers"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

func ChangeRole(c *gin.Context) {
	var reqBody struct {
		ID   uuid.UUID            `json:"id" binding:"required"`
		Role repository.UserRoles `json:"role" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		controllers.HandleValidations(c, err)
		return
	}

	user, err := db.Repo.UpdateUserRole(c, repository.UpdateUserRoleParams{
		ID:   reqBody.ID,
		Role: reqBody.Role,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
