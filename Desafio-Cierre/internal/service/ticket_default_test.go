package service_test

import (
	"context"
	"goweb/Desafio-Cierre/internal"
	"goweb/Desafio-Cierre/internal/handler"
	"goweb/Desafio-Cierre/internal/repository"
	"goweb/Desafio-Cierre/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		// arrange
		// - repository:
		rp := repository.NewRepositoryTicketMap(map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "USA",
				Hour:    "10:00",
				Price:   100,
			},
			2: {
				Name:    "Michael",
				Email:   "michael@gmail.com",
				Country: "Japan",
				Hour:    "10:00",
				Price:   100,
			},
		}, 2)

		// - service
		sv := service.NewServiceTicketDefault(rp)

		country := "USA"

		hd := handler.NewDefaultTicketsHandler(sv)
		req := &http.Request{}
		chiCtx := chi.NewRouteContext()
		chiCtx.URLParams.Add("dest", country)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))

		res := httptest.NewRecorder()
		// act
		hdFunc := hd.GetTicketsAmount()

		hdFunc(res, req)

		// assert
		expectedCode := http.StatusOK
		expectedBody := `{"data":1,"message":"success"}`
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
		require.Equal(t, expectedHeader, res.Header())
	})
}
