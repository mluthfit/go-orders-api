package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RequiredResponse struct {
	code    int
	message string
}

func NewRequiredResponse() *RequiredResponse {
	return &RequiredResponse{code: 200}
}

func resError(ctx *gin.Context, res *RequiredResponse, err any) {
	ctx.AbortWithStatusJSON(res.code, gin.H{
		"code":    res.code,
		"errors":  err,
		"message": res.message,
	})
}

func resSuccess(ctx *gin.Context, res *RequiredResponse, data any) {
	ctx.JSON(res.code, gin.H{
		"code":    res.code,
		"data":    data,
		"message": res.message,
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
