package example

import (
	"context"
	"log"
	"os"

	swo "github.com/solarwinds/swo-client-go/pkg/client"
)

func Setup() (context.Context, *swo.Client) {
	baseUrl := getEnvVar("SWO_BASE_URL")
	apiToken := getEnvVar("SWO_API_TOKEN")

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
