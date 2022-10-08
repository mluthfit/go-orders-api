package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func resError(ctx *gin.Context, code int, err any) {
	var message = "request validation errors"
	if value, ok := err.(string); ok {
		message = value
		err = nil
	}

	ctx.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
		"errors":  err,
	})
}

func resSuccess(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, gin.H{
		"code": code,
		"data": data,
	})
}

func extractBindError(err error) map[string][]string {
	var errors = map[string][]string{}

	for _, value := range err.(validator.ValidationErrors) {
		var errMessage = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", value.Field(), value.ActualTag())

		if _, isExist := errors[value.Field()]; isExist {
			errors[value.Field()] = append(errors[value.Field()], errMessage)
			continue
		}

		errors[value.Field()] = []string{errMessage}
	}

	return errors
}
