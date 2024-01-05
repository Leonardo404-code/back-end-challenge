package main

import (
	"github.com/gin-gonic/gin"
)

// @contact.name Leonardo Bispo
// @title Frete Rápido back-end challenge
// @version 1.0
// @description Frete Rápido challenge to develop Rest API for external queries and return only expected values.
func main() {
	r := gin.Default()

	if err := r.Run(":" + "3000"); err != nil {
		panic(err)
	}
}
