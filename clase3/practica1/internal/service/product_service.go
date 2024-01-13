package service

import (
	"fmt"
	"goweb/clase3/practica1/internal"
	"time"
)

func NewProductDefaultService(rp internal.ProductRepository) *ProductDefaultService {
	return &ProductDefaultService{
		rp: rp,
	}
}

type ProductDefaultService struct {
	rp internal.ProductRepository
}

func (p *ProductDefaultService) AddNewProduct(product *internal.Product) (err error) {
	//validations, not empty types allowed
	if validateWrongFields(*product) {
		fmt.Println("Product invalid field")
		return internal.ErrProductInvalidField
	}

	if checkDateFormat(*product) {
		fmt.Println("Product invalid date format")
		return internal.ErrProductInvalidDate
	}
	(*p).rp.AddNewProduct(product)
	return err
}
func (p *ProductDefaultService) GetProductById(id int) (product internal.Product, err error) {
	return (*p).rp.GetProductById(id)
}

func (p *ProductDefaultService) UpdateProduct(product *internal.Product) (err error) {
	if validateWrongFields(*product) {
		fmt.Println("Product invalid field")
		return internal.ErrProductInvalidField
	}

	if checkDateFormat(*product) {
		fmt.Println("Product invalid date format")
		return internal.ErrProductInvalidDate
	}
	err = p.rp.UpdateProduct(product)
	return err
}

func (p *ProductDefaultService) DeleteProduct(id int) (err error) {
	return (*p).rp.DeleteProduct(id)
}

func checkDateFormat(p internal.Product) bool {
	var format string = "2006-01-02 15:04:05"
	fmt.Println(p)
	_, err := time.Parse(format, p.Expiration+" 00:00:00")

	if err != nil {

		fmt.Println("TimeParse Invalid Date Format", err)
		return true
	}

	return false
}

func validateWrongFields(p internal.Product) bool {
	// Acá también podría pasarle un map con los bytes que leo de la request para validad directamente la request y no el dato parseado
	return (p.CodeValue == "" || p.Name == "" ||
		p.Expiration == "" || p.Price == 0.0 ||
		p.Quantity == 0)
}
