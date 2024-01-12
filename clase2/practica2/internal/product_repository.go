package internal

type ProductRepository interface {
	AddNewProduct(p Product) error
	GetProductById(id int) (Product, error)
}
