package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

var (
	websiteMockData = struct {
		fieldName string
		fieldUrl  string
	}{
		"swo-client-go - website",
		"www.solarwinds.com",
	}
)

func GetCreateInput() CreateWebsiteInput {
	inputJson, err := ioutil.ReadFile("website_test.json")
	if err != nil {
		log.Fatal(err)
	}

	var website CreateWebsiteInput
	if err = json.Unmarshal(inputJson, &website); err != nil {
		log.Fatal(err)
	}

	return website
}

func TestSwoService_ReadWebsite(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__getWebsiteByIdInput](r)
		if err != nil {
			t.Errorf("Swo.ReadWebsite error: %v", err)
		}

		got := gqlInput.Id
		want := "123"

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		// response := &getWebsiteByIdEntitiesEntityQueriesByIdWebsite{
		// 	Name: &website.Name,
		// 	Url:  website.Url,
		// 	Monitoring: &getWebsiteByIdEntitiesEntityQueriesByIdWebsiteMonitoring{
		// 		Availability: &getWebsiteByIdEntitiesEntityQueriesByIdWebsiteMonitoringAvailabilityAvailabilityMonitoring{
		// 			Protocols: website.AvailabilityCheckSettings.Protocols,
		// 		},
		// 		CustomHeaders: []getWebsiteByIdEntitiesEntityQueriesByIdWebsiteMonitoringCustomHeadersCustomHeader{
		// 			{
		// 				Name:  website.AvailabilityCheckSettings.CustomHeaders[0].Name,
		// 				Value: website.AvailabilityCheckSettings.CustomHeaders[0].Value,
		// 			},
		// 		},
		// 		Rum: &getWebsiteByIdEntitiesEntityQueriesByIdWebsiteMonitoringRumRumMonitoring{
		// 			ApdexTimeInSeconds: &website.AvailabilityCheckSettings.TestIntervalInSeconds,
		// 			Spa:                *website.Rum.Spa,
		// 		},
		// 	},
		// }

		sendGraphQLResponse(t, w, getWebsiteByIdResponse{
			Entities: getWebsiteByIdEntitiesEntityQueries{},
		})
	})

	got, err := client.WebsiteService().Read(ctx, "123")
	if err != nil {
		t.Errorf("Swo.ReadWebsite returned error: %v", err)
	}

	want := &getWebsiteByIdResponse{
		Entities: getWebsiteByIdEntitiesEntityQueries{},
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.ReadNotification returned %+v, wanted %+v", got, want)
	}
}

func TestSwoService_CreateWebsite(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	id := uuid.NewString()
	input := GetCreateInput()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createWebsiteMutationInput](r)
		if err != nil {
			t.Errorf("Swo.CreateWebsite error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, createWebsiteMutationResponse{
			Dem: createWebsiteMutationDemDemMutations{
				CreateWebsite: createWebsiteMutationDemDemMutationsCreateWebsiteCreateWebsiteSuccess{
					Id: id,
				},
			},
		})
	})

	got, err := client.WebsiteService().Create(ctx, input)
	if err != nil {
		t.Errorf("Swo.CreateWebsite returned error: %v", err)
	}

	want := &CreateWebsiteResult{
		Id: id,
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.CreateNotification returned %+v, want %+v", got, want)
	}
}

func TestSwoService_UpdateWebsite(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := UpdateWebsiteInput{
		Id: "123",
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__updateWebsiteMutationInput](r)
		if err != nil {
			t.Errorf("Swo.UpdateWebsite error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, updateWebsiteMutationResponse{
			Dem: updateWebsiteMutationDemDemMutations{
				UpdateWebsite: updateWebsiteMutationDemDemMutationsUpdateWebsiteUpdateWebsiteSuccess{
					Id: input.Id,
				},
			},
		})
	})

	if err := client.WebsiteService().Update(ctx, input); err != nil {
		t.Errorf("Swo.UpdateWebsite returned error: %v", err)
	}
}

func TestSwoService_DeleteWebsite(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := DeleteWebsiteInput{
		Id: "123",
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__deleteWebsiteMutationInput](r)
		if err != nil {
			t.Errorf("Swo.DeleteWebsite error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Swo.DeleteWebsite: Request got = %+v, want %+v", got, want)
		}

		sendGraphQLResponse(t, w, deleteWebsiteMutationResponse{
			Dem: deleteWebsiteMutationDemDemMutations{
				DeleteWebsite: deleteWebsiteMutationDemDemMutationsDeleteWebsiteDeleteWebsiteSuccess{
					Id: input.Id,
				},
			},
		})
	})

	err := client.WebsiteService().Delete(ctx, input.Id)
	if err != nil {
		t.Errorf("Swo.DeleteWebsite returned error: %v", err)
	}
}

func TestSwoService_WebsiteServerErrors(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", httpErrorResponse)

	_, err := client.WebsiteService().Create(ctx, CreateWebsiteInput{})
	if err == nil {
		t.Error("Swo.WebsiteServerErrors expected an error response")
	}
	_, err = client.WebsiteService().Read(ctx, "123")
	if err == nil {
		t.Error("Swo.WebsiteServerErrors expected an error response")
	}
	err = client.WebsiteService().Update(ctx, UpdateWebsiteInput{})
	if err == nil {
		t.Error("Swo.WebsiteServerErrors expected an error response")
	}
	err = client.WebsiteService().Delete(ctx, "123")
	if err == nil {
		t.Error("Swo.WebsiteServerErrors expected an error response")
	}
}

// func TestWebsite_Marshal(t *testing.T) {
// 	testJSONMarshal(t, &ReadWebsiteResult{}, "{}")

// 	id := uuid.NewString()
// 	got := ReadWebsiteResult{
// 		Entities: getWebsiteByIdEntitiesEntityQueries{
// 			ById: &getWebsiteByIdEntitiesEntityQueriesByIdEntity{},
// 		},
// 	}

// 	want := fmt.Sprintf(`
// 	{
// 		"id": "%s",
// 		"type": "email",
// 		"title": "%s",
// 		"description": "%s",
// 		"settings": {
// 			"addresses": [
// 				{
// 					"email": "test1@host.com"
// 				},
// 				{
// 					"email": "test2@host.com"
// 				}
// 			]
// 		},
// 		"createdAt": "%s",
// 		"createdBy": "%s"
// 	}`,
// 		id,
// 		notificationsMockData.fieldTitle,
// 		notificationsMockData.fieldDesc,
// 		notificationsMockData.fieldCreatedAt.Format(time.RFC3339),
// 		notificationsMockData.fieldCreatedBy)

// 	testJSONMarshal(t, got, want)
// }
