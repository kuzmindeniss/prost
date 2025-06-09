package units

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

	applications, err := db.Repo.GetApplicationsByUnitID(c, reqBody.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные: " + err.Error(),
		})
		return
	}

	if len(applications) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Невозможно удалить подразделение, так как оно используется в заявках",
		})
		return
	}

	err = db.Repo.DeleteUnit(c, reqBody.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
