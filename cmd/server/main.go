package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"shipping-calculator-api/internal/database"
	shippingHandlers "shipping-calculator-api/pkg/shipping/handler"
)

// @contact.name Leonardo Bispo
// @title Shipping Calculator API
// @version 1.0
// @description A REST API that fetches delivery data from the Frete Rápido API, stores it in the database and calculates the shipping data.
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
