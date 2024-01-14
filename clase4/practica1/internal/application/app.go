package application

import (
	"goweb/clase4/practica1/internal"
	"goweb/clase4/practica1/internal/handler"
	"goweb/clase4/practica1/internal/repository"
	"goweb/clase4/practica1/internal/service"
	"goweb/clase4/practica1/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewDefaultHTTP creates a new instance of a default http server
func NewDefaultHTTP(addr string) *DefaultHTTP {
	// default config / values
	// ...

	return &DefaultHTTP{
		addr: addr,
	}
}

// DefaultHTTP is a struct that represents the default implementation of a http server
type DefaultHTTP struct {
	// addr is the address of the http server
	addr string
}

// Run runs the http server
func (h *DefaultHTTP) Run() (err error) {
	// initialize dependencies
	// - repository
	productsMap := make(map[int]internal.Product)
	storage.ReadProducts(&productsMap, "./storage/intial_data.json")
	rp := repository.NewProductMap(productsMap, 0)
	// - service
	sv := service.NewProductDefaultService(rp)
	// - handler
	hd := handler.NewDefaultProducts(sv)
	// - router
	rt := chi.NewRouter()
	//   endpoints
	rt.Post("/products", hd.AddNewProductHandler())
	rt.Get("/products/{id}", hd.GetProductByIdHandler())
	rt.Put("/products/{id}", hd.UpdateProductHandler())
	rt.Patch("/products/{id}", hd.UpdatePartialHandler())
	rt.Delete("/products/{id}", hd.DeleteProductHandler())
	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}
