package errors

import (
	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, statusCode int, err error) {
	errMessage := map[string]string{
		"error": err.Error(),
	}
	ctx.JSON(statusCode, errMessage)
}
