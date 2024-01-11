package handler

import (
	"fmt"
	"goweb/clase2/practica2/internal"
	"goweb/clase2/practica2/internal/repository"
	"goweb/clase2/practica2/internal/service"
	"goweb/clase2/practica2/platform/web/request"
	"goweb/clase2/practica2/platform/web/response"
	"net/http"
)

func AddNewProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct internal.Product
		if err := request.JSON(r, &newProduct); err != nil {
			fmt.Println(err)
		}

		fmt.Println(newProduct)

		if service.ValidateWrongFields(newProduct) {
			fmt.Println("Product invalid field")
			response.JSON(w, http.StatusBadRequest, "Product invalid field")
			return
		}

		if service.CheckUniqueCode(newProduct) {
			fmt.Println("Product code value already exist")
			response.JSON(w, http.StatusBadRequest, "Product code value already exist")
			return
		}

		if service.CheckDateFormat(newProduct) {
			fmt.Println("Product invalid date format")
			response.JSON(w, http.StatusBadRequest, "product invalid date format")
			return
		}

		newProduct.Id = getNewId()
		repository.AddNewProduct(newProduct)
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product created",
			"data":    newProduct,
		})
	}
}

func getNewId() int {
	return repository.Products[len(repository.Products)-1].Id + 1
}
