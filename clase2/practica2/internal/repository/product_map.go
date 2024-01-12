package repository

import (
	"errors"
	"goweb/clase2/practica2/internal"
)

type ProductMap struct {
	products map[int]internal.Product
	lastId   int
}

func NewProductMap(products map[int]internal.Product, lastId int) *ProductMap {
	return &ProductMap{
		products: products,
		lastId:   lastId,
	}
}

func (m *ProductMap) GetProductById(id int) (product internal.Product, err error) {
	for _, p := range (*m).products {
		if p.Id == id {
			product = p
			break
		}
	}
	return product, errors.New("Product not found")
}

func (m *ProductMap) AddNewProduct(product internal.Product) (err error) {
	for _, p := range (*m).products {
		if p.CodeValue == product.CodeValue {
			return errors.New("Product code value already exist")
		}
	}
	(*m).lastId++
	product.Id = (*m).lastId
	m.products[product.Id] = product
	return err
}
