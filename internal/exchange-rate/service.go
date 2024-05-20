package exchangerate

import (
	"fmt"

	"github.com/2marks/go-exchangerate-api/types"
)

type Service struct {
	repo  types.ExchangeRateRepository
	cache types.Cache
}

var SUPPORTED_FROM_CURRENCIES = []string{"EUR"}

func NewService(repo types.ExchangeRateRepository, cache types.Cache) *Service {
	return &Service{repo: repo, cache: cache}
}

func (s *Service) GetExchangeRate(fromCurrency string, toCurrency string) (any, error) {
	if !s.IsFromCurrencySupported(fromCurrency) {
		return nil, fmt.Errorf("from currency (%s) not supported", fromCurrency)
	}

	cacheKey := fmt.Sprintf("EXCHANGE_RATE:%s-%s", fromCurrency, toCurrency)
	cachedExchangeRates := new(types.ExchangeRateApiSuccessResponse)
	_, err := s.cache.Get(cacheKey, cachedExchangeRates)
	if err == nil {
		return s.reformExchangeRatesData(cachedExchangeRates), nil
	}

	exchangeRates, err := s.repo.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return nil, err
	}

	s.cache.Set(cacheKey, *exchangeRates)

	return s.reformExchangeRatesData(exchangeRates), err
}

func (s *Service) IsFromCurrencySupported(currency string) bool {
	for _, cur := range SUPPORTED_FROM_CURRENCIES {
		if cur == currency {
			return true
		}
	}

	return false
}

func (s *Service) reformExchangeRatesData(data *types.ExchangeRateApiSuccessResponse) types.ExchangeRateResponse {
	return types.ExchangeRateResponse{
		From:        data.Base,
		Rates:       data.Rates,
		GeneratedAt: data.Date,
	}
}
