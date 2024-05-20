package main

import (
	"fmt"
	"log"

	"github.com/2marks/go-exchangerate-api/cmd/api"
	"github.com/2marks/go-exchangerate-api/config"
)

func main() {
	apiServer := api.NewApiServer(fmt.Sprintf(":%s", config.Envs.Port))

	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
	}
}
