package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Поле обязательно"
	case "lte":
		return "Должно быть меньше чем " + fe.Param()
	case "gte":
		return "Должно быть больше чем " + fe.Param()
	}
	return "Неизвестная ошибка"
}

func HandleValidations(context *gin.Context, err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
		}
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		return out
	}

	return nil
}
