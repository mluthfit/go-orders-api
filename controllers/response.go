package controllers

import "github.com/gin-gonic/gin"

func resError(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, gin.H{
		"code":   code,
		"errors": err.Error(),
	})
}

func resSuccess(ctx *gin.Context, code int, data any) {
	ctx.JSON(code, gin.H{
		"code": code,
		"data": data,
	})
}
