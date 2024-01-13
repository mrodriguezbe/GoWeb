package application

import (
	"goweb/clase3/practica1/internal"
	"goweb/clase3/practica1/internal/handler"
	"goweb/clase3/practica1/internal/repository"
	"goweb/clase3/practica1/internal/service"
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
	rt.Post("/products", hd.AddNewProductHandler())
	rt.Get("/products/{id}", hd.GetProductByIdHandler())
	rt.Put("/products/{id}", hd.UpdateProductHandler())
	rt.Patch("/products/{id}", hd.UpdatePartialHandler())
	rt.Delete("/products/{id}", hd.DeleteProductHandler())

	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}
