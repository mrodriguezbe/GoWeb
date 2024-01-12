package service

import (
	"goweb/clase2/practica2/internal"
)

func NewProductDefaultService(rp internal.ProductRepository) *ProductDefaultService {
	return &ProductDefaultService{
		rp: rp,
	}
}

type ProductDefaultService struct {
	rp internal.ProductRepository
}

func (p *ProductDefaultService) AddNewProduct(product internal.Product) (err error) {
	//validations, not empty types allowed
	//..
	(*p).rp.AddNewProduct(product)
	return err
}
func (p *ProductDefaultService) GetProductById(id int) (product internal.Product, err error) {
	//validations, not empty types allowed
	//..
	return (*p).rp.GetProductById(id)
}
