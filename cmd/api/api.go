package api

import (
	"fmt"
	"net/http"

	"github.com/2marks/go-exchangerate-api/cache"
	exchangerate "github.com/2marks/go-exchangerate-api/internal/exchange-rate"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{addr: addr}
}

func (a *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	/* init cache */
	redisClient := cache.GetRedisClient()
	cache := cache.NewCache(redisClient)
	/* end init cache */

	/*start exchange rate routes*/
	exchangeRateRepository := &exchangerate.Respository{}
	exchangeRateService := exchangerate.NewService(exchangeRateRepository, cache)
	exchangeRateHandler := exchangerate.NewHandler(exchangeRateService)
	exchangeRateHandler.RegisterRoutes(subRouter)
	/*end exchange rate routes*/

	fmt.Println("server listening on", a.addr)

	return http.ListenAndServe(a.addr, router)
}
