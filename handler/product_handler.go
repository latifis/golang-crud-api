package handler

import (
	"net/http"

	"golang-crud-api/model"
	"golang-crud-api/service"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := service.GetProducts()
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product model.Product

	if err := c.BindJSON(&product); err != nil {
		return
	}

	service.CreateProduct(product)

	c.JSON(http.StatusCreated, product)
}
