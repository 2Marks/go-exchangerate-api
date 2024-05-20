package types

type ExchangeRateApiSuccessResponse struct {
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

type ExchangeRateApiErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

type ExchangeRateService interface {
	GetExchangeRate(fromCurrency string, toCurrency string) (any, error)
}

type ExchangeRateRepository interface {
	GetExchangeRate(fromCurrency string, toCurrency string) (*ExchangeRateApiSuccessResponse, error)
}

type GetExchangeRateDTO struct {
	FromCurrency string `json:"fromCurrency"`
	ToCurrency   string `json:"toCurrency"`
}

type ExchangeRateResponse struct {
	From        string             `json:"from"`
	Rates       map[string]float64 `json:"rates"`
	GeneratedAt string             `json:"generated_at"`
}
