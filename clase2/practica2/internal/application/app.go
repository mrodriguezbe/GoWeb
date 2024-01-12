package application

import (
	"goweb/clase2/practica2/internal"
	"goweb/clase2/practica2/internal/handler"
	"goweb/clase2/practica2/internal/repository"
	"goweb/clase2/practica2/internal/service"
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
	rp := repository.NewProductMap(make(map[int]internal.Product), 0)
	// - service
	sv := service.NewProductDefaultService(rp)
	// - handler
	hd := handler.NewDefaultProducts(sv)
	// - router
	rt := chi.NewRouter()
	//   endpoints
	rt.Post("/products", hd.GetProductByIdHandler())
	rt.Get("/products/{id}", hd.AddNewProductHandler())

	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}
