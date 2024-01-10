package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	Name       string  `json:"name"`
	Id         int     `json:"id"`
	Quantity   int     `json:"quantity"`
	CodeValue  string  `json:"code_value"`
	Published  bool    `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
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

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func addNewProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newProduct)

	if validateWrongFields(newProduct) {
		fmt.Println("Product invalid field")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Product Invalid Field"))
		return
	}

	if checkUniqueCode(newProduct) {
		fmt.Println("Product code value already exist")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Product code value already existd"))
		return
	}

	if checkDateFormat(newProduct) {
		fmt.Println("Product invalid date format")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Product invalid date format"))
		return
	}

	newProduct.Id = getNewId()
	Products = append(Products, newProduct)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New product added ID:" + strconv.Itoa(newProduct.Id)))
}

func getNewId() int {
	return Products[len(Products)-1].Id + 1
}

func checkDateFormat(p Product) bool {
	var format string = "24/02/1997"
	_, err := time.Parse(format, p.Expiration)

	if err != nil {
		fmt.Println("TimeParse Invalid Date Format")
		return false
	}

	return true
}

func checkUniqueCode(p Product) bool {
	for _, e := range Products {
		if e.CodeValue == p.CodeValue {
			return true
		}
	}
	return false
}

func validateWrongFields(p Product) bool {
	return (p.CodeValue == "" || p.Name == "" ||
		p.Expiration == "" || p.Price == 0.0 ||
		p.Quantity == 0)
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

func main() {
	ReadProducts()
	r := chi.NewRouter()
	r.Get("/ping", pingHandler)
	r.Post("/product", addNewProductHandler)
	r.Get("/products/{id}", getProductByIdHandler)

	fmt.Println("El servidor está escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
