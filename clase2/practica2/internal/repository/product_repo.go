package repository

import "goweb/clase2/practica2/internal"

var Products []internal.Product

func NewProductRepo() []internal.Product {
	return Products
}

func GetProductById(id int) (prod internal.Product) {
	for _, p := range Products {
		if p.Id == id {
			prod = p
			break
		}
	}
	return prod
}

func AddNewProduct(p internal.Product) {
	Products = append(Products, p)
}
