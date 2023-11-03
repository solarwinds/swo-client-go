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

	Read(ctx, client, create.Id)
	Update(ctx, client, create.Id)
}

func GetCreateInput() swo.CreateWebsiteInput {
	inputJson, err := ioutil.ReadFile(createFile)
	if err != nil {
		log.Fatal(err)
	}

	var website swo.CreateWebsiteInput
	if err = json.Unmarshal(inputJson, &website); err != nil {
		log.Fatal(err)
	}

	return website
}

func Create(ctx context.Context, client *swo.Client) *swo.CreateWebsiteResult {
	website := GetCreateInput()

	result, err := client.WebsiteService().Create(ctx, website)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Read(ctx context.Context, client *swo.Client, id string) *swo.ReadWebsiteResult {
	result, err := client.WebsiteService().Read(ctx, id)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func Update(ctx context.Context, client *swo.Client, id string) {
	inputJson, err := ioutil.ReadFile(createFile)
	if err != nil {
		log.Fatal(err)
	}

	var input swo.UpdateWebsiteInput
	if err = json.Unmarshal(inputJson, &input); err != nil {
		log.Fatal(err)
	}

	input.Id = id
	input.Name += "->[UPDATED_WEBSITE]"

	if err = client.WebsiteService().Update(ctx, input); err != nil {
		log.Fatal(err)
	}
}

func Delete(ctx context.Context, client *swo.Client, id string) {
	if err := client.WebsiteService().Delete(ctx, id); err != nil {
		log.Fatal(err)
	}
}
