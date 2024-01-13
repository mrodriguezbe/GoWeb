package repository

import (
	"goweb/clase3/practica2/internal"
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
			return product, err
		}
	}
	return product, internal.ErrProductNotFound
}

func (m *ProductMap) AddNewProduct(product *internal.Product) (err error) {
	for _, p := range (*m).products {
		if p.CodeValue == product.CodeValue {
			return internal.ErrProductCodeValueAlreadyExists
		}
	}
	(*m).lastId++
	product.Id = (*m).lastId
	m.products[product.Id] = *product
	return err
}

func (m *ProductMap) UpdateProduct(product *internal.Product) (err error) {
	if _, err := m.GetProductById(product.Id); err != nil {
		return internal.ErrProductNotFound
	}

	m.products[product.Id] = *product
	return
}

func (m *ProductMap) DeleteProduct(id int) (err error) {
	_, ok := m.products[id]

	if !ok {
		err = internal.ErrProductNotFound
		return
	}

	delete(m.products, id)
	return
}
