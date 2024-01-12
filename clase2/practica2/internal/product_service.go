package internal

type ProductService interface {
	AddNewProduct(p Product) (err error)
	GetProductById(id int) (product Product, err error)
}
