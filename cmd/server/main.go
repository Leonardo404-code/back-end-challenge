package main

import (
	"github.com/gin-gonic/gin"

	"frete-rapido-api/internal/database"
	"frete-rapido-api/internal/env"
)

// @contact.name Leonardo Bispo
// @title Frete Rápido back-end challenge
// @version 1.0
// @description Frete Rápido challenge to develop Rest API for external queries and return only expected values.
func main() {
	_ = database.Must()

	r := gin.Default()
	if err := r.Run(":" + env.GetString("PORT")); err != nil {
		panic(err)
	}
}
