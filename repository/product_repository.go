package repository

import "golang-crud-api/model"

var products []model.Product

func GetAll() []model.Product {
	return products
}

func Create(product model.Product) {
	products = append(products, product)
}
