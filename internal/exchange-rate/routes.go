package exchangerate

import (
	"net/http"

	"github.com/2marks/go-exchangerate-api/types"
	"github.com/2marks/go-exchangerate-api/utils"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Handler struct {
	exchangeRateService types.ExchangeRateService
}

func NewHandler(service types.ExchangeRateService) *Handler {
	return &Handler{exchangeRateService: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/exchange-rates", h.GetExchangeRate).Methods(http.MethodGet)
}

func (h *Handler) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	var payload = new(types.GetExchangeRateDTO)
	var decoder = schema.NewDecoder()

	if err := decoder.Decode(payload, r.URL.Query()); err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	data, err := h.exchangeRateService.GetExchangeRate(payload.FromCurrency, payload.ToCurrency)
	if err != nil {
		utils.WriteErrorToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessToJson(w, http.StatusOK, "exchange rate fetched successfully", data)
}
