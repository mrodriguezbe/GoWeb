package internal

import "errors"

var (
	// ErrMovieTitleAlreadyExists is the error returned when a movie title already exists
	ErrProductCodeValueAlreadyExists = errors.New("product code value already exists")
	// ErrMovieNotFound is the error returned when a movie is not found
	ErrProductNotFound = errors.New("product not found")
)

type ProductRepository interface {
	AddNewProduct(product *Product) (err error)
	GetProductById(id int) (product Product, err error)
	UpdateProduct(prodruct *Product) (err error)
	DeleteProduct(id int) (err error)
}
