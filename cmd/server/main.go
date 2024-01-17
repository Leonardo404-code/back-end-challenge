package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"shipping-calculator-api/internal/database"
	shippingHandlers "shipping-calculator-api/pkg/shipping/handler"
)

// @contact.name Leonardo Bispo
// @title Frete Rápido back-end challenge
// @version 1.0
// @description Frete Rápido challenge to develop Rest API for external queries and return only expected values.
func main() {
	dbConnection := database.Must()

	r := gin.Default()
	shipHandler := shippingHandlers.Build(dbConnection)

	r.POST("/quote", shipHandler.Quotes)
	r.GET("/metrics", shipHandler.Metrics)

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
