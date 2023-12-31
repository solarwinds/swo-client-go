package client

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	notificationsMockData = struct {
		fieldTitle     string
		fieldDesc      string
		fieldCreatedAt time.Time
		fieldCreatedBy string
		emailSettings  map[string]any
	}{
		"swo-client-go - title",
		"swo-client-go - description",
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"123456789",
		map[string]any{
			"addresses": []any{
				map[string]any{"email": string("test1@host.com")},
				map[string]any{"email": string("test2@host.com")},
			},
		},
	}
)

func TestSwoService_ReadNotification(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	var settings any = notificationsMockData.emailSettings

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__getNotificationInput](r)
		if err != nil {
			t.Errorf("Swo.ReadNotification error: %v", err)
		}

		sendGraphQLResponse(t, w, getNotificationResponse{
			User: getNotificationUserAuthenticatedUser{
				CurrentOrganization: getNotificationUserAuthenticatedUserCurrentOrganization{
					NotificationServiceConfiguration: ReadNotificationResult{
						Id:          gqlInput.ConfigurationId,
						Type:        gqlInput.ConfigurationType,
						Title:       notificationsMockData.fieldTitle,
						Description: &notificationsMockData.fieldDesc,
						Settings:    &settings,
						CreatedAt:   notificationsMockData.fieldCreatedAt,
						CreatedBy:   notificationsMockData.fieldCreatedBy,
					},
				},
			},
		})
	})

	got, err := client.NotificationsService().Read(ctx, "123", "email")
	if err != nil {
		t.Errorf("Swo.ReadNotification returned error: %v", err)
	}

	want := &ReadNotificationResult{
		Id:          "123",
		Title:       notificationsMockData.fieldTitle,
		Description: &notificationsMockData.fieldDesc,
		Type:        "email",
		Settings:    &settings,
		CreatedAt:   notificationsMockData.fieldCreatedAt,
		CreatedBy:   notificationsMockData.fieldCreatedBy,
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.ReadNotification returned %+v, wanted %+v", got, want)
	}
}

func TestSwoService_CreateNotification(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	var settings any = notificationsMockData.emailSettings

	id := uuid.NewString()
	input := CreateNotificationInput{
		Title:       notificationsMockData.fieldTitle,
		Description: &notificationsMockData.fieldDesc,
		Type:        "email",
		Settings:    settings,
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createNotificationInput](r)
		if err != nil {
			t.Errorf("Swo.CreateNotification error: %v", err)
		}

		got := gqlInput.Configuration
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, createNotificationResponse{
			CreateNotificationServiceConfiguration: createNotificationCreateNotificationServiceConfigurationCreateNotificationServiceConfigurationResponse{
				Code:    "201",
				Success: true,
				Message: "",
				Configuration: &CreateNotificationResult{
					Id:          id,
					Type:        got.Type,
					Title:       got.Title,
					Description: got.Description,
					Settings:    &got.Settings,
					CreatedAt:   notificationsMockData.fieldCreatedAt,
					CreatedBy:   notificationsMockData.fieldCreatedBy,
				},
			},
		})
	})

	got, err := client.NotificationsService().Create(ctx, input)
	if err != nil {
		t.Errorf("Swo.CreateNotification returned error: %v", err)
	}

	want := &CreateNotificationResult{
		Id:          id,
		Title:       input.Title,
		Description: input.Description,
		Type:        input.Type,
		Settings:    &input.Settings,
		CreatedAt:   notificationsMockData.fieldCreatedAt,
		CreatedBy:   notificationsMockData.fieldCreatedBy,
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.CreateNotification returned %+v, want %+v", got, want)
	}
}

func TestSwoService_UpdateNotification(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	var settings any = notificationsMockData.emailSettings

	input := UpdateNotificationInput{
		Id:          "123",
		Title:       &notificationsMockData.fieldTitle,
		Description: &notificationsMockData.fieldDesc,
		Settings:    &settings,
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__updateNotificationInput](r)
		if err != nil {
			t.Errorf("Swo.UpdateNotification error: %v", err)
		}

		got := gqlInput.Configuration
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, updateNotificationResponse{
			UpdateNotificationServiceConfiguration: &updateNotificationUpdateNotificationServiceConfigurationUpdateNotificationServiceConfigurationResponse{
				Code:    "201",
				Success: true,
				Message: "",
				Configuration: &updateNotificationUpdateNotificationServiceConfigurationUpdateNotificationServiceConfigurationResponseConfigurationNotificationService{
					Id:          got.Id,
					Title:       *got.Title,
					Description: got.Description,
					Settings:    got.Settings,
				},
			},
		})
	})

	_, err := client.NotificationsService().Update(ctx, input)
	if err != nil {
		t.Errorf("Swo.UpdateNotification returned error: %v", err)
	}
}

func TestSwoService_DeleteNotification(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := DeleteNotificationServiceConfigurationInput{
		Id: "123",
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__deleteNotificationInput](r)
		if err != nil {
			t.Errorf("Swo.DeleteNotification error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Swo.DeleteNotification: Request got = %+v, want %+v", got, want)
		}

		sendGraphQLResponse(t, w, deleteNotificationResponse{
			DeleteNotificationServiceConfiguration: &deleteNotificationDeleteNotificationServiceConfigurationDeleteNotificationServiceConfigurationResponse{
				Code:    "201",
				Success: true,
				Message: "",
			},
		})
	})

	err := client.NotificationsService().Delete(ctx, input.Id)
	if err != nil {
		t.Errorf("Swo.DeleteNotification returned error: %v", err)
	}
}

func TestSwoService_NotificationsServerErrors(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", httpErrorResponse)

	_, err := client.NotificationsService().Create(ctx, CreateNotificationInput{})
	if err == nil {
		t.Error("Swo.NotificationsServerErrors expected an error response")
	}
	_, err = client.NotificationsService().Read(ctx, "123", "email")
	if err == nil {
		t.Error("Swo.NotificationsServerErrors expected an error response")
	}
	_, err = client.NotificationsService().Update(ctx, UpdateNotificationInput{})
	if err == nil {
		t.Error("Swo.NotificationsServerErrors expected an error response")
	}
	err = client.NotificationsService().Delete(ctx, "123")
	if err == nil {
		t.Error("Swo.NotificationsServerErrors expected an error response")
	}
}

func TestNotification_Marshal(t *testing.T) {
	testJSONMarshal(t, &ReadNotificationResult{}, "{}")

	var settings any = notificationsMockData.emailSettings
	id := uuid.NewString()

	got := ReadNotificationResult{
		Id:          id,
		Title:       notificationsMockData.fieldTitle,
		Description: &notificationsMockData.fieldDesc,
		Type:        "email",
		Settings:    &settings,
		CreatedAt:   notificationsMockData.fieldCreatedAt,
		CreatedBy:   notificationsMockData.fieldCreatedBy,
	}

	want := fmt.Sprintf(`
	{
		"id": "%s",
		"type": "email",
		"title": "%s",
		"description": "%s",
		"settings": {
			"addresses": [
				{
					"email": "test1@host.com"
				},
				{
					"email": "test2@host.com"
				}
			]
		},
		"createdAt": "%s",
		"createdBy": "%s"
	}`,
		id,
		notificationsMockData.fieldTitle,
		notificationsMockData.fieldDesc,
		notificationsMockData.fieldCreatedAt.Format(time.RFC3339),
		notificationsMockData.fieldCreatedBy)

	testJSONMarshal(t, got, want)
}
