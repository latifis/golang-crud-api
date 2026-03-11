package main

import (
	"golang-crud-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/products", handler.GetProducts)
	router.POST("/products", handler.CreateProduct)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
