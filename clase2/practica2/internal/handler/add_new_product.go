package handler

import (
	"fmt"
	"goweb/clase2/practica2/internal"
	"goweb/clase2/practica2/platform/web/request"
	"goweb/clase2/practica2/platform/web/response"
	"net/http"
	"time"
)

func NewDefaultProducts(sv internal.ProductService) *DefaultProductService {
	//defualt config / values

	return &DefaultProductService{
		sv: sv,
	}
}

type DefaultProductService struct {
	sv internal.ProductService
}

type BodyProductJSON struct {
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	CodeValue  string  `json:"code_value"`
	Published  bool    `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}

func (d *DefaultProductService) AddNewProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProductJSON BodyProductJSON
		if err := request.JSON(r, &newProductJSON); err != nil {
			fmt.Println(err)
		}

		product := internal.Product{
			Name:       newProductJSON.Name,
			Quantity:   newProductJSON.Quantity,
			CodeValue:  newProductJSON.CodeValue,
			Published:  newProductJSON.Published,
			Expiration: newProductJSON.Expiration,
			Price:      newProductJSON.Price,
		}

		if validateWrongFields(product) {
			fmt.Println("Product invalid field")
			response.JSON(w, http.StatusBadRequest, "Product invalid field")
			return
		}

		if checkDateFormat(product) {
			fmt.Println("Product invalid date format")
			response.JSON(w, http.StatusBadRequest, "product invalid date format")
			return
		}

		err := (*d).sv.AddNewProduct(product)
		if err != nil {
			fmt.Println(err)
			response.JSON(w, http.StatusBadRequest, "Product code value already exist")
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product created",
			"data":    product,
		})
	}
}

func checkDateFormat(p internal.Product) bool {
	var format string = "24/02/1997"
	_, err := time.Parse(format, p.Expiration)

	if err != nil {
		fmt.Println("TimeParse Invalid Date Format")
		return false
	}

	return true
}

func validateWrongFields(p internal.Product) bool {
	return (p.CodeValue == "" || p.Name == "" ||
		p.Expiration == "" || p.Price == 0.0 ||
		p.Quantity == 0)
}
