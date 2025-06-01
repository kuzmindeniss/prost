package units

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/bff/helpers"
)

func GetAll(c *gin.Context) {
	units, err := helpers.Repo.GetUnits(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при получении подразделений: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"units": units,
	})
}
