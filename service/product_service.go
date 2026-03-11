package service

import (
	"golang-crud-api/model"
	"golang-crud-api/repository"
)

func GetProducts() []model.Product {
	return repository.GetAll()
}

func CreateProduct(product model.Product) {
	repository.Create(product)
}
