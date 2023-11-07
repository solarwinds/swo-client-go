package main

import (
	"context"
	"log"

	"github.com/solarwindscloud/swo-client-go/example"
	swo "github.com/solarwindscloud/swo-client-go/pkg/client"
)

const (
	createFile = "create.json"

	test commit
)

func main() {
	ctx, client := example.Setup()

	uri := Create(ctx, client)
	defer Delete(ctx, client, uri.Id)

	Read(ctx, client, uri.Id)
	Update(ctx, client, uri.Id)
}

func Create(ctx context.Context, client *swo.Client) *swo.CreateUriResult {
	input, err := swo.GetObjectFromFile[swo.CreateUriInput](createFile)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.UriService().Create(ctx, *input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swo.Client, id string) *swo.ReadUriResult {
	result, err := client.UriService().Read(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swo.Client, id string) {
	input, err := swo.GetObjectFromFile[swo.UpdateUriInput](createFile)
	if err != nil {
		log.Fatal(err)
	}

	input.Id = id
	input.Name += "->[UPDATED_URI]"

	if err = client.UriService().Update(ctx, *input); err != nil {
		log.Fatal(err)
	}
}

func Delete(ctx context.Context, client *swo.Client, id string) {
	if err := client.UriService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
