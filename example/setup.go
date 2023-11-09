package example

import (
	"context"
	"log"
	"os"

	swo "github.com/solarwindscloud/swo-client-go/pkg/client"
)

var (
	apiTokenVar = "SWO_API_TOKEN"
	baseUrlVar  = "SWO_BASE_URL"
)

func Setup() (context.Context, *swo.Client) {
	baseUrl := getEnvVar(baseUrlVar)
	apiToken := getEnvVar(apiTokenVar)

	ctx := context.Background()

	client, err := swo.New(apiToken,
		swo.BaseUrlOption(baseUrl),
		swo.DebugOption(true),
	)

	if err != nil {
		log.Fatal(err)
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
