package errors

import (
	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, statusCode int, err error) {
	_ = ctx.Error(err)
	ctx.Status(statusCode)
}
