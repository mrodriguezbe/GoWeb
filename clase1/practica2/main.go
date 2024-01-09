package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	Id         int     `json:id`
	Name       string  `json:name`
	Quantity   int     `json:quantity`
	CodeValue  string  `json:code_value`
	Published  bool    `json:is_published`
	Expiration string  `json:expiration`
	Price      float64 `json:price`
}

var Products []Product

func ReadProducts() {

	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(data, &Products); err != nil {
		fmt.Println("Error de unmarshal", err)
	}
}

func getProductsMoreExpensiveThanHandler(w http.ResponseWriter, r *http.Request) {
	var price float64
	var err error
	price, err = strconv.ParseFloat(r.URL.Query().Get("Price"), 64)

	if err != nil {
		fmt.Println(err)
	}

	var response []Product
	for _, p := range Products {
		if p.Price <= price {
			response = append(response, p)
		}
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println(err)
	}
}

func getProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi((chi.URLParam(r, "id")))

	if err != nil {
		fmt.Println("Error de atoi", err)
	}

	for _, p := range Products {
		if p.Id == id {
			err := json.NewEncoder(w).Encode(p)
			if err != nil {
				fmt.Println("Error de encoding", err)
			}
			break
		}
	}
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(Products)
	if err != nil {
		fmt.Println(err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	ReadProducts()
	r := chi.NewRouter()
	r.Get("/ping", pingHandler)
	r.Get("/products", getProductsHandler)
	r.Get("/products/{id}", getProductByIdHandler)
	r.Get("/products/search", getProductsMoreExpensiveThanHandler)

	http.ListenAndServe(":8080", r)
	fmt.Println("El servidor está escuchando en http://localhost:8080")
}
