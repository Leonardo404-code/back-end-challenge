package shipping

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Quotes(ctx *gin.Context)
	Metrics(ctx *gin.Context)
}
