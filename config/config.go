package config

import (
	"os"
)

type (
	Config struct {
		APIGatewayAddress string `json:"api_gateway_address"`
	}
)

var Instance = Config{}

func Read() (config Config) {
	config = readFromEnvVar()

	return
}

func readFromEnvVar() (config Config) {
	config.APIGatewayAddress = readEnvVarWithDefaultValue("API_GATEWAY_ADDRESS", "http://localhost")

	return
}

func readEnvVarWithDefaultValue(key, defaultValue string) string {
	if envVarValue, ok := os.LookupEnv(key); ok {
		return envVarValue
	}
	return defaultValue
}
