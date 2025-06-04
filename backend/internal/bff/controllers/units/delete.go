package units

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kuzmindeniss/prost/internal/bff/controllers"
	"github.com/kuzmindeniss/prost/internal/bff/helpers"
)

func Delete(c *gin.Context) {
	var reqBody struct {
		ID uuid.UUID `json:"id" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		controllers.HandleValidations(c, err)
		return
	}

	err := helpers.Repo.DeleteUnit(c, reqBody.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
