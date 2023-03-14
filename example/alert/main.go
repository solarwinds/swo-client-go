package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
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
	Update(ctx, client, read.Id)
}

func Create(ctx context.Context, client *swo.Client) *swo.CreateAlertDefinitionResult {
	inputJson, err := ioutil.ReadFile(createFile)
	if err != nil {
		log.Fatal(err)
	}

	var input swo.AlertDefinitionInput
	if err = json.Unmarshal(inputJson, &input); err != nil {
		log.Fatal(err)
	}

	result, err := client.AlertsService().Create(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swo.Client, id string) *swo.ReadAlertDefinitionResult {
	result, err := client.AlertsService().Read(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swo.Client, id string) *swo.UpdateAlertDefinitionResult {
	inputJson, err := ioutil.ReadFile(createFile)
	if err != nil {
		log.Fatal(err)
	}

	var input swo.AlertDefinitionInput
	if err = json.Unmarshal(inputJson, &input); err != nil {
		log.Fatal(err)
	}

	input.Name += "->[UPDATE_ALERT]"

	result, err := client.AlertsService().Update(ctx, id, input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Delete(ctx context.Context, client *swo.Client, id string) {
	if err := client.AlertsService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
