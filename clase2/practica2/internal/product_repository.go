package internal

type ProductRepository interface {
	AddNewProduct(product Product) (err error)
	GetProductById(id int) (product Product, err error)
}
