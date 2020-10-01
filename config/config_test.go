package config

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	setDummyEnvVar()
	expectedOutput := getDummyConfig()

	config := Read()

	if !reflect.DeepEqual(config, expectedOutput) {
		t.Fatal("unexpected output")
	}
}

func setDummyEnvVar() {
	dummyConfig := getDummyConfig()

	os.Setenv("API_GATEWAY_ADDRESS", dummyConfig.APIGatewayAddress)
}

func getDummyConfig() Config {
	dummyConfig := Config{
		APIGatewayAddress: "http://localhost",
	}

	return dummyConfig
}
