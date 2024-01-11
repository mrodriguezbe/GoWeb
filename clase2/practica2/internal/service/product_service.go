package service

import (
	"encoding/json"
	"fmt"
	"goweb/clase2/practica2/internal"
	"goweb/clase2/practica2/internal/repository"
	"os"
	"time"
)

func ReadProducts() {

	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(data, &repository.Products); err != nil {
		fmt.Println("Error de unmarshal", err)
	}
}

func CheckDateFormat(p internal.Product) bool {
	var format string = "24/02/1997"
	_, err := time.Parse(format, p.Expiration)

	if err != nil {
		fmt.Println("TimeParse Invalid Date Format")
		return false
	}

	return true
}

func CheckUniqueCode(p internal.Product) bool {
	for _, e := range repository.Products {
		if e.CodeValue == p.CodeValue {
			return true
		}
	}
	return false
}

func ValidateWrongFields(p internal.Product) bool {
	return (p.CodeValue == "" || p.Name == "" ||
		p.Expiration == "" || p.Price == 0.0 ||
		p.Quantity == 0)
}
