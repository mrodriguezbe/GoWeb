package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GreetingRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	var request GreetingRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	greeting := fmt.Sprintf("Hello %s %s", request.FirstName, request.LastName)
	w.Write([]byte(greeting))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("pong"))
}

func main() {
	r := chi.NewRouter()
	r.Post("/greetings", greetingsHandler)
	r.Get("/ping", pingHandler)

	http.ListenAndServe(":8080", r)
	fmt.Println("El servidor está escuchando en http://localhost:8080")
}
