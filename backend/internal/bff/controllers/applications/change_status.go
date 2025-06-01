package applications

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kuzmindeniss/prost/internal/bff/controllers"
	"github.com/kuzmindeniss/prost/internal/bff/helpers"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

func ChangeStatus(c *gin.Context) {
	var reqBody struct {
		ID     uuid.UUID                    `json:"id" binding:"required"`
		Status repository.ApplicationStatus `json:"status" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		controllers.HandleValidations(c, err)
		return
	}

	application, err := helpers.Repo.UpdateApplicationStatus(c, repository.UpdateApplicationStatusParams{
		ID:     reqBody.ID,
		Status: reqBody.Status,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"application": application,
	})
}
