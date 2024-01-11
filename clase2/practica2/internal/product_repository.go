package internal

type ProductRepository interface {
	AddNewProduct(p Product)
	GetProductById(id int) Product
}
