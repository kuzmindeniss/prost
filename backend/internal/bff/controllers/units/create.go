package units

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/bff/controllers"
	"github.com/kuzmindeniss/prost/internal/bff/helpers"
)

func Create(c *gin.Context) {
	var reqBody struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		controllers.HandleValidations(c, err)
		return
	}

	unit, err := helpers.Repo.CreateUnit(c, reqBody.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при создании подразделения: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"unit": unit,
	})
}
