package exchangerate

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/2marks/go-exchangerate-api/config"
	"github.com/2marks/go-exchangerate-api/types"
)

type Respository struct{}

func (r *Respository) GetExchangeRate(fromCurrency string, toCurrency string) (*types.ExchangeRateApiSuccessResponse, error) {
	var apiKey = config.Envs.ExchangeRateApiKey
	var url = fmt.Sprintf("http://exchangesrateapi.com/api/latest?apiKey=%s&base=%s&symbols=%s", apiKey, fromCurrency, toCurrency)

	fmt.Printf("about to call exchange rate third-party api. from:%s, to:%s", fromCurrency, toCurrency)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		errorResponse := new(types.ExchangeRateApiErrorResponse)
		if err := json.NewDecoder(res.Body).Decode(errorResponse); err != nil {
			return nil, err
		}

		log.Println(errorResponse.Message)

		return nil, fmt.Errorf("error occured from exchange rate api: %s", errorResponse.Message)
	}

	var responseBody = new(types.ExchangeRateApiSuccessResponse)
	if err := json.NewDecoder(res.Body).Decode(responseBody); err != nil {
		return nil, fmt.Errorf("error occured while parsing response from exchange rate api %s", err.Error())
	}

	return responseBody, nil
}
