package applications

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kuzmindeniss/prost/internal/db"
)

func GetAll(c *gin.Context) {
	applications, err := db.Repo.GetApplications(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при получении заявок: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"applications": applications,
	})
}
