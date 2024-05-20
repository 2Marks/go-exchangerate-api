package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	ExchangeRateApiKey  string
	ExchangeRateBaseUrl string
	RedisUrl            string
	RedisPassword       string
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment variables from file")
	}

	return Config{
		Port:                getEnvStr("PORT", ""),
		ExchangeRateBaseUrl: getEnvStr("EXCHANGE_RATE_BASE_URL", ""),
		ExchangeRateApiKey:  getEnvStr("EXCHANGE_RATE_API_KEY", ""),
		RedisUrl:            getEnvStr("REDIS_URL", "localhost:6379"),
		RedisPassword:       getEnvStr("REDIS_PASSWORD", ""),
	}
}

func getEnvStr(key string, fallback string) string {
	value, ok := os.LookupEnv(key)

	if !ok && fallback == "" {
		log.Fatalf("env variable %s not found", key)
	}

	if !ok && fallback != "" {
		return fallback
	}

	return value
}
