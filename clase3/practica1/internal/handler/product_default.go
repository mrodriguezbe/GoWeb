package handler

import (
	"fmt"
	"goweb/clase3/practica1/internal"
	"goweb/clase3/practica1/platform/web/request"
	"goweb/clase3/practica1/platform/web/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	Id         int     `json:"id"`
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
			response.Text(w, http.StatusBadRequest, "Invalid Body")
			return
		}

		product := internal.Product{
			Name:       newProductJSON.Name,
			Quantity:   newProductJSON.Quantity,
			CodeValue:  newProductJSON.CodeValue,
			Published:  newProductJSON.Published,
			Expiration: newProductJSON.Expiration,
			Price:      newProductJSON.Price,
		}

		err := (*d).sv.AddNewProduct(&product)
		if err != nil {
			fmt.Println(err)
			response.JSON(w, http.StatusBadRequest, "Product code value already exist/Invalid Field")
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product created",
			"data":    product,
		})
	}
}

func (d *DefaultProductService) GetProductByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi((chi.URLParam(r, "id")))

		if err != nil {
			fmt.Println("Error de atoi", err)
			return
		}

		product, err := (*d).sv.GetProductById(id)

		if err != nil {
			fmt.Println(err)
			response.JSON(w, http.StatusNotFound, "Product not found")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product found",
			"data":    product,
		})
	}
}

func (d *DefaultProductService) UpdateProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi((chi.URLParam(r, "id")))

		if err != nil {
			fmt.Println("Error de atoi", err)
			response.Text(w, http.StatusBadRequest, "Invalid Id")
			return
		}

		var productJSON BodyProductJSON
		if err := request.JSON(r, &productJSON); err != nil {
			fmt.Println(err)
			response.Text(w, http.StatusBadRequest, "Invalid Body")
			return
		}

		product := internal.Product{
			Id:         id,
			Name:       productJSON.Name,
			Quantity:   productJSON.Quantity,
			CodeValue:  productJSON.CodeValue,
			Published:  productJSON.Published,
			Expiration: productJSON.Expiration,
			Price:      productJSON.Price,
		}

		if err := d.sv.UpdateProduct(&product); err != nil {
			fmt.Println(err)
			response.Error(w, http.StatusNotFound, "Product not found")
			return
		}

		data := BodyProductJSON{
			Id:         id,
			Name:       product.Name,
			Quantity:   product.Quantity,
			CodeValue:  product.CodeValue,
			Published:  product.Published,
			Expiration: product.Expiration,
			Price:      product.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})

	}
}

func (d *DefaultProductService) UpdatePartialHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi((chi.URLParam(r, "id")))

		if err != nil {
			fmt.Println("Error de atoi", err)
			response.Text(w, http.StatusBadRequest, "Invalid Id")
			return
		}

		product, err := d.sv.GetProductById(id)

		data := BodyProductJSON{
			Id:         id,
			Name:       product.Name,
			Quantity:   product.Quantity,
			CodeValue:  product.CodeValue,
			Published:  product.Published,
			Expiration: product.Expiration,
			Price:      product.Price,
		}
		if err := request.JSON(r, &data); err != nil {
			fmt.Println(err)
			response.Text(w, http.StatusBadRequest, "Invalid Body")
		}

		product = internal.Product{
			Id:         id,
			Name:       data.Name,
			Quantity:   data.Quantity,
			CodeValue:  data.CodeValue,
			Published:  data.Published,
			Expiration: data.Expiration,
			Price:      data.Price,
		}

		if err := d.sv.UpdateProduct(&product); err != nil {
			fmt.Println(err)
			response.Error(w, http.StatusNotFound, "Product not found")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product updated",
			"data":    data,
		})
	}
}

func (d *DefaultProductService) DeleteProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi((chi.URLParam(r, "id")))

		if err != nil {
			fmt.Println("Error de atoi", err)
			response.Text(w, http.StatusBadRequest, "Invalid Id")
			return
		}

		if err := d.sv.DeleteProduct(id); err != nil {
			fmt.Println(err)
			response.Error(w, http.StatusNotFound, "Product not found")
			return
		}

		response.JSON(w, http.StatusNoContent, "movie deleted")
	}
}
