package units

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kuzmindeniss/prost/internal/bff/controllers"
	"github.com/kuzmindeniss/prost/internal/db"
	"github.com/kuzmindeniss/prost/internal/db/repository"
)

func ChangeName(c *gin.Context) {
	var reqBody struct {
		ID   uuid.UUID `json:"id" binding:"required"`
		Name string    `json:"name" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		controllers.HandleValidations(c, err)
		return
	}

	unit, err := db.Repo.UpdateUnitName(c, repository.UpdateUnitNameParams{
		ID:   reqBody.ID,
		Name: reqBody.Name,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"unit": unit,
	})
}
