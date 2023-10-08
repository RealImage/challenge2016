package main

import (
	"github.com/gin-gonic/gin"
)

func SetResponse(ctx *gin.Context, httpStatusCode int, data interface{}) {
	res := map[string]interface{} {
		"data"	: data,
	}
	ctx.JSON(httpStatusCode, res)
}

func LogError(err interface{}) {
	panic(err)
}