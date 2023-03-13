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

	read := Read(ctx, client, create.Id, create.Type)
	Update(ctx, client, *read)
}

func Create(ctx context.Context, client *swoClient.Client) *swoClient.CreateNotificationResult {
	inputJson, err := ioutil.ReadFile(createFile)
	if err != nil {
		log.Fatal(err)
	}

	var input swoClient.CreateNotificationInput
	if err = json.Unmarshal(inputJson, &input); err != nil {
		log.Fatal(err)
	}

	result, err := client.NotificationsService().Create(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swoClient.Client, id string, notificationType string) *swoClient.ReadNotificationResult {
	result, err := client.NotificationsService().Read(ctx, id, notificationType)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swoClient.Client, notification swoClient.ReadNotificationResult) {
	input, err := swoClient.ConvertObject[swoClient.UpdateNotificationInput](notification)
	if err != nil {
		log.Fatal(err)
	}

	input.Title = swoClient.Ptr(*input.Title + "->[UPDATE_NOTIFICATION]")

	_, err = client.NotificationsService().Update(ctx, *input)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(ctx context.Context, client *swoClient.Client, id string) {
	if err := client.NotificationsService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
