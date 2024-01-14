package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"goweb/clase4/practica1/internal"
	"goweb/clase4/practica1/internal/handler"
	"os"
)

func ReadProducts(productsMap *map[int]internal.Product, path string) (err error) {

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		err = errors.New("File does not exist")
		return
	}
	defer f.Close()

	jsonSlice := make([]handler.BodyProductJSON, 0)
	err = json.NewDecoder(f).Decode(&jsonSlice)

	if err != nil {
		fmt.Println("Decoder error", err)
		err = errors.New("Decoder error")
		return
	}

	for i, v := range jsonSlice {
		(*productsMap)[i] = internal.Product{
			Id:         i,
			Name:       v.Name,
			Quantity:   v.Quantity,
			CodeValue:  v.CodeValue,
			Published:  v.Published,
			Expiration: v.Expiration,
			Price:      v.Price,
		}
	}
	return
}

func WriteProducts(productsMap *map[int]internal.Product, path string) (err error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		fmt.Println(err)
		err = errors.New("Cannot open file")
		return
	}

	jsonMap := make([]handler.BodyProductJSON, 0)
	for k, v := range *productsMap {
		jsonMap = append(jsonMap, handler.BodyProductJSON{
			Id:         k,
			Name:       v.Name,
			Quantity:   v.Quantity,
			CodeValue:  v.CodeValue,
			Published:  v.Published,
			Expiration: v.Expiration,
			Price:      v.Price,
		})
	}

	err = json.NewEncoder(f).Encode(jsonMap)

	if err != nil {
		fmt.Println("Encoder error", err)
		err = errors.New("Encoder error")
		return
	}
	return
}
