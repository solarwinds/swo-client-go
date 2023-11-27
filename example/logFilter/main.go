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

	logFilter := Create(ctx, client)
	defer Delete(ctx, client, logFilter.Id)

	Read(ctx, client, logFilter.Id)
	Update(ctx, client, logFilter.Id)
}

func Create(ctx context.Context, client *swo.Client) *swo.CreateLogFilterResult {
	input, err := swo.GetObjectFromFile[swo.CreateExclusionFilterInput](createFile)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.LogFilterService().Create(ctx, *input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swo.Client, id string) *swo.ReadLogFilterResult {
	result, err := client.LogFilterService().Read(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swo.Client, id string) {
	input, err := swo.GetObjectFromFile[swo.UpdateExclusionFilterInput](createFile)
	if err != nil {
		log.Fatal(err)
	}

	input.Id = id
	input.Name += "->[UPDATED_URI]"

	if err = client.LogFilterService().Update(ctx, *input); err != nil {
		log.Fatal(err)
	}
}

func Delete(ctx context.Context, client *swo.Client, id string) {
	if err := client.LogFilterService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
