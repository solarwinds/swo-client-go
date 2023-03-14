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

	read := Read(ctx, client, create.Id, create.Type)
	Update(ctx, client, *read)
}

func Create(ctx context.Context, client *swo.Client) *swo.CreateNotificationResult {
	inputJson, err := ioutil.ReadFile(createFile)
	if err != nil {
		log.Fatal(err)
	}

	var input swo.CreateNotificationInput
	if err = json.Unmarshal(inputJson, &input); err != nil {
		log.Fatal(err)
	}

	result, err := client.NotificationsService().Create(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swo.Client, id string, notificationType string) *swo.ReadNotificationResult {
	result, err := client.NotificationsService().Read(ctx, id, notificationType)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swo.Client, notification swo.ReadNotificationResult) {
	input, err := swo.ConvertObject[swo.UpdateNotificationInput](notification)
	if err != nil {
		log.Fatal(err)
	}

	input.Title = swo.Ptr(*input.Title + "->[UPDATE_NOTIFICATION]")

	_, err = client.NotificationsService().Update(ctx, *input)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(ctx context.Context, client *swo.Client, id string) {
	if err := client.NotificationsService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
