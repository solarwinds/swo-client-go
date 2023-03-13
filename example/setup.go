package example

import (
	"context"
	"log"
	"os"

	swoClient "github.com/solarwindscloud/terraform-provider-swo/internal/client"
)

var (
	apiTokenVar = "SWO_API_TOKEN"
	baseUrlVar  = "SWO_BASE_URL"
)

func Setup() (context.Context, *swoClient.Client) {
	baseUrl := getEnvVar(baseUrlVar)
	apiToken := getEnvVar(apiTokenVar)
	ctx := context.Background()
	client := swoClient.NewClient(apiToken,
		swoClient.BaseUrlOption(baseUrl),
		swoClient.DebugOption(true),
	)

	if client == nil {
		log.Fatal("Unable to create an instance of the SWO client.")
		return nil, nil
	}

	return ctx, client
}

func getEnvVar(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("Missing %s environment variable or variable is not set.", name)
	}

	return value
}
