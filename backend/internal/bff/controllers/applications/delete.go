package applications

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kuzmindeniss/prost/internal/bff/controllers"
	"github.com/kuzmindeniss/prost/internal/db"
)

func Delete(c *gin.Context) {
	var reqBody struct {
		ID uuid.UUID `json:"id" binding:"required"`
	}

	if err := c.BindJSON(&reqBody); err != nil {
		controllers.HandleValidations(c, err)
		return
	}

	err := db.Repo.DeleteApplication(c, reqBody.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
