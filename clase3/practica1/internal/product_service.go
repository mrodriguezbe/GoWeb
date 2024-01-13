package internal

import "errors"

var (
	ErrProductInvalidField = errors.New("Product invalid field")
	ErrProductInvalidDate  = errors.New("Product invalid date format")
)

type ProductService interface {
	AddNewProduct(product *Product) (err error)
	GetProductById(id int) (product Product, err error)
	UpdateProduct(product *Product) (err error)
	DeleteProduct(id int) (err error)
}
