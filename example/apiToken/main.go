package main

import (
	"context"
	"log"

	"github.com/solarwinds/swo-client-go/example"
	swo "github.com/solarwinds/swo-client-go/pkg/client"
)

const (
	createFile = "create.json"
)

func main() {
	ctx, client := example.Setup()

	apiToken := Create(ctx, client)
	defer Delete(ctx, client, apiToken.Id)

	Read(ctx, client, apiToken.Id)
	Update(ctx, client, apiToken.Id)
}

func Create(ctx context.Context, client *swo.Client) *swo.CreateApiTokenResult {
	input, err := swo.GetObjectFromFile[swo.CreateTokenInput](createFile)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.ApiTokenService().Create(ctx, *input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swo.Client, id string) *swo.ReadApiTokenResult {
	result, err := client.ApiTokenService().Read(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swo.Client, id string) {
	input, err := swo.GetObjectFromFile[swo.UpdateTokenInput](createFile)
	if err != nil {
		log.Fatal(err)
	}

	input.Id = id
	*input.Name += "->[UPDATED_TOKEN]"

	if err = client.ApiTokenService().Update(ctx, *input); err != nil {
		log.Fatal(err)
	}
}

func Delete(ctx context.Context, client *swo.Client, id string) {
	if err := client.ApiTokenService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
