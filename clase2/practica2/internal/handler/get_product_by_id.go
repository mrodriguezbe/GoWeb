package handler

import (
	"fmt"
	"goweb/clase2/practica2/platform/web/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (d *DefaultProductService) GetProductByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi((chi.URLParam(r, "id")))

		if err != nil {
			fmt.Println("Error de atoi", err)
		}

		product, err := (*d).sv.GetProductById(id)

		if err != nil {
			fmt.Println(err)
			response.JSON(w, http.StatusNotFound, "Product not found")
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "product found",
			"data":    product,
		})
	}
}
