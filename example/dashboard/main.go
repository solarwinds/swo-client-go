package main

import (
	"context"
	"log"

	"github.com/solarwindscloud/swo-client-go/example"
	swo "github.com/solarwindscloud/swo-client-go/pkg/client"
)

const (
	createFile = "create.json"
)

func main() {
	ctx, client := example.Setup()

	create := Create(ctx, client)
	defer Delete(ctx, client, create.Id)

	read := Read(ctx, client, create.Id)
	Update(ctx, client, *read)
}

func Create(ctx context.Context, client *swo.Client) *swo.CreateDashboardResult {
	input, err := swo.GetObjectFromFile[swo.CreateDashboardInput](createFile)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.DashboardsService().Create(ctx, *input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swo.Client, id string) *swo.ReadDashboardResult {
	result, err := client.DashboardsService().Read(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swo.Client, dashboard swo.ReadDashboardResult) *swo.UpdateDashboardResult {
	input, err := swo.ConvertObject[swo.UpdateDashboardInput](dashboard)
	if err != nil {
		log.Fatal(err)
	}

	input.Name += "->[UPDATE_DASHBOARD]"

	result, err := client.DashboardsService().Update(ctx, *input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Delete(ctx context.Context, client *swo.Client, id string) {
	if err := client.DashboardsService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
