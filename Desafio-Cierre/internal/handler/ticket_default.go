package handler

import (
	"goweb/Desafio-Cierre/internal"
	"goweb/Desafio-Cierre/platform/web/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewDefaultTicketsHandler(sv internal.ServiceTicket) *DefaultTicketsHandler {
	//defualt config / values

	return &DefaultTicketsHandler{
		sv: sv,
	}
}

type DefaultTicketsHandler struct {
	sv internal.ServiceTicket
}

// Get the proportion of people traveling to a country
func (h *DefaultTicketsHandler) GetProportion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest := chi.URLParam(r, "dest")

		result, err := h.sv.GetProportionByDestination(dest)

		if err != nil {
			response.Error(w, http.StatusNotFound, "Destination not found")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    result,
		})
	}
}

// Get the total amount of tickets that travel to a country
func (h *DefaultTicketsHandler) GetTicketsAmount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest := chi.URLParam(r, "dest")

		result, err := h.sv.GetTicketsByDestination(dest)

		if err != nil {
			response.Error(w, http.StatusNotFound, "Destination not found")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    result,
		})
	}
}
