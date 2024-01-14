package handler_test

import (
	"context"
	"fmt"
	"goweb/clase4/practica1/internal"
	"goweb/clase4/practica1/internal/handler"
	"goweb/clase4/practica1/internal/repository"
	"goweb/clase4/practica1/internal/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGetById(t *testing.T) {
	t.Run("Success get", func(t *testing.T) {
		db := map[int]internal.Product{
			1: {
				Id:         1,
				Name:       "Te",
				Quantity:   1,
				CodeValue:  "testcodevalue",
				Published:  true,
				Expiration: "dd/mm/yyyy",
				Price:      50.0,
			},
		}

		rp := repository.NewProductMap(db, 1)
		sv := service.NewProductDefaultService(rp)
		hd := handler.NewDefaultProducts(sv)
		hdFunc := hd.GetProductByIdHandler()

		req := &http.Request{}
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		res := httptest.NewRecorder()

		hdFunc(res, req)

		expectedCode := http.StatusOK
		expectedBody := `{"data":{"Id":1,"Name":"Te","Quantity":1,"CodeValue":"testcodevalue","Published":true,"Expiration":"dd/mm/yyyy","Price":50},"message":"product found"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})

	t.Run("Failure get", func(t *testing.T) {
		db := map[int]internal.Product{
			1: {
				Id:         1,
				Name:       "Te",
				Quantity:   1,
				CodeValue:  "testcodevalue",
				Published:  true,
				Expiration: "dd/mm/yyyy",
				Price:      50.0,
			},
		}

		rp := repository.NewProductMap(db, 1)
		sv := service.NewProductDefaultService(rp)
		hd := handler.NewDefaultProducts(sv)
		hdFunc := hd.GetProductByIdHandler()

		req := &http.Request{}
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "2")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		res := httptest.NewRecorder()

		hdFunc(res, req)

		expectedCode := http.StatusNotFound
		expectedBody := "Product not found"
		expectedHeader := http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}}

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}

func TestPost(t *testing.T) {
	t.Run("Success Post", func(t *testing.T) {
		db := make(map[int]internal.Product)

		rp := repository.NewProductMap(db, 0)
		sv := service.NewProductDefaultService(rp)
		hd := handler.NewDefaultProducts(sv)
		hdFunc := hd.AddNewProductHandler()

		req := &http.Request{
			Body: io.NopCloser(strings.NewReader(fmt.Sprintf(
				`{"name": "Te", "quantity": 1, "code_value": "testcodevalue", "is_published": true, "expiration": "1996-03-15", "price": 50}`),
			)),
			Header: http.Header{"Content-Type": []string{"application/json"}},
		}

		res := httptest.NewRecorder()

		hdFunc(res, req)

		expectedCode := http.StatusCreated
		expectedBody := `{"data":{"Id":1,"Name":"Te","Quantity":1,"CodeValue":"testcodevalue","Published":true,"Expiration":"1996-03-15","Price":50},"message":"Product created"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})
}

func TestDelete(t *testing.T) {
	t.Run("Success Delete", func(t *testing.T) {
		db := map[int]internal.Product{
			1: {
				Id:         1,
				Name:       "Te",
				Quantity:   1,
				CodeValue:  "testcodevalue",
				Published:  true,
				Expiration: "dd/mm/yyyy",
				Price:      50.0,
			},
		}

		rp := repository.NewProductMap(db, 0)
		sv := service.NewProductDefaultService(rp)
		hd := handler.NewDefaultProducts(sv)
		hdFunc := hd.DeleteProductHandler()

		req := &http.Request{}
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		res := httptest.NewRecorder()

		hdFunc(res, req)

		expectedCode := http.StatusNoContent
		expectedBody := "Product deleted"
		expectedHeader := http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}}

		require.Equal(t, expectedCode, res.Code)
		require.Equal(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())

	})
}
