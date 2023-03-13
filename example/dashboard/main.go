package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/solarwindscloud/terraform-provider-swo/example"
	swoClient "github.com/solarwindscloud/terraform-provider-swo/internal/client"
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

func Create(ctx context.Context, client *swoClient.Client) *swoClient.CreateDashboardResult {
	inputJson, err := ioutil.ReadFile(createFile)
	if err != nil {
		log.Fatal(err)
	}

	var input swoClient.CreateDashboardInput
	if err = json.Unmarshal(inputJson, &input); err != nil {
		log.Fatal(err)
	}

	result, err := client.DashboardsService().Create(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swoClient.Client, id string) *swoClient.ReadDashboardResult {
	result, err := client.DashboardsService().Read(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swoClient.Client, dashboard swoClient.ReadDashboardResult) *swoClient.UpdateDashboardResult {
	input, err := swoClient.ConvertObject[swoClient.UpdateDashboardInput](dashboard)
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

func Delete(ctx context.Context, client *swoClient.Client, id string) {
	if err := client.DashboardsService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
