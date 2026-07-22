package client

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	dashboardsMockData = struct {
		fieldName       string
		fieldUpdatedAt  time.Time
		fieldOwnerId    string
		fieldOwnerName  string
		fieldCategoryId string
		versionId       int
	}{
		"swo-client-go - title",
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		"123456789",
		"owner name",
		"123",
		2,
	}
)

func TestService_CreateDashboard(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	id := uuid.NewString()

	input := CreateDashboardInput{
		Name:       dashboardsMockData.fieldName,
		IsPrivate:  Ptr(true),
		CategoryId: &dashboardsMockData.fieldCategoryId,
		Layout: []LayoutInput{
			{Id: "123", X: 0, Y: 0, Height: 2, Width: 2},
		},
		Widgets: []WidgetInput{
			{Id: "123", Type: "Proportional", Properties: nil},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createDashboardInput](r)
		if err != nil {
			t.Errorf("Swo.CreateDashboard error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, createDashboardResponse{
			CreateDashboard: createDashboardCreateDashboardCreateDashboardResponse{
				Code:    "200",
				Success: true,
				Message: "",
				Dashboard: &createDashboardCreateDashboardCreateDashboardResponseDashboard{
					Id:        id,
					OwnerId:   &dashboardsMockData.fieldOwnerId,
					CreatedAt: dashboardsMockData.fieldUpdatedAt,
					UpdatedAt: dashboardsMockData.fieldUpdatedAt,
					Version:   &dashboardsMockData.versionId,
				},
			},
		})
	})

	got, err := client.DashboardsService().Create(ctx, input)
	if err != nil {
		t.Errorf("Swo.CreateDashboard returned error: %v", err)
	}

	want := &CreateDashboardResult{
		Id:        id,
		OwnerId:   &dashboardsMockData.fieldOwnerId,
		CreatedAt: dashboardsMockData.fieldUpdatedAt,
		UpdatedAt: dashboardsMockData.fieldUpdatedAt,
		Version:   &dashboardsMockData.versionId,
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.CreateDashboard returned %+v, want %+v", got, want)
	}
}

func TestService_ReadDashboard(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__getDashboardByIdInput](r)
		if err != nil {
			t.Errorf("Swo.ReadDashboard error: %v", err)
		}

		sendGraphQLResponse(t, w, getDashboardByIdResponse{
			Dashboards: &getDashboardByIdDashboardsDashboardQueries{
				ByIdOrSystemReference: &getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboard{
					Id:        gqlInput.Id,
					Name:      dashboardsMockData.fieldName,
					IsPrivate: Ptr(true),
					UpdatedAt: dashboardsMockData.fieldUpdatedAt,
					CreatedAt: dashboardsMockData.fieldUpdatedAt,
					Category: &getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardCategory{
						Id: dashboardsMockData.fieldCategoryId,
					},
					OwnerId: &dashboardsMockData.fieldOwnerId,
					Layout: []getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardLayout{
						{Id: "123", X: 0, Y: 0, Height: 2, Width: 2},
					},
					Widgets: []getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardWidgetsWidget{
						{Id: "123", Type: "Proportional"},
					},
				},
			},
		})
	})

	id := uuid.NewString()
	got, err := client.DashboardsService().Read(ctx, id)
	if err != nil {
		t.Errorf("Swo.ReadDashboard returned error: %v", err)
	}

	want := &ReadDashboardResult{
		Id:        id,
		Name:      dashboardsMockData.fieldName,
		IsPrivate: Ptr(true),
		UpdatedAt: dashboardsMockData.fieldUpdatedAt,
		CreatedAt: dashboardsMockData.fieldUpdatedAt,
		Category: &getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardCategory{
			Id: dashboardsMockData.fieldCategoryId,
		},
		OwnerId: &dashboardsMockData.fieldOwnerId,
		Layout: []getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardLayout{
			{Id: "123", X: 0, Y: 0, Height: 2, Width: 2},
		},
		Widgets: []getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardWidgetsWidget{
			{Id: "123", Type: "Proportional", Properties: nil},
		},
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.ReadDashboard returned %+v, wanted %+v", got, want)
	}
}

func TestService_UpdateDashboard(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	versionId := 2
	input := UpdateDashboardInput{
		Id:         "123",
		Name:       dashboardsMockData.fieldName,
		CategoryId: &dashboardsMockData.fieldCategoryId,
		Layout: []LayoutInput{
			{Id: "123", X: 0, Y: 0, Height: 2, Width: 2},
		},
		Widgets: []WidgetInput{
			{Id: "123", Type: "Proportional"},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__updateDashboardInput](r)
		if err != nil {
			t.Errorf("Swo.UpdateDashboard error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, updateDashboardResponse{
			UpdateDashboard: updateDashboardUpdateDashboardUpdateDashboardResponse{
				Code:    "200",
				Success: true,
				Message: "",
				Dashboard: &updateDashboardUpdateDashboardUpdateDashboardResponseDashboard{
					Id:      got.Id,
					Version: &versionId,
				},
			},
		})
	})

	_, err := client.DashboardsService().Update(ctx, input)
	if err != nil {
		t.Errorf("Swo.UpdateDashboard returned error: %v", err)
	}
}

func TestService_DeleteDashboard(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := DeleteDashboardInput{Id: "123"}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__deleteDashboardInput](r)
		if err != nil {
			t.Errorf("Swo.DeleteDashboard error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Swo.DeleteDashboard: Request got = %+v, want %+v", got, want)
		}

		sendGraphQLResponse(t, w, deleteDashboardResponse{
			DeleteDashboard: deleteDashboardDeleteDashboardDeleteDashboardResponse{
				Success: true,
			},
		})
	})

	err := client.DashboardsService().Delete(ctx, input.Id)
	if err != nil {
		t.Errorf("Swo.DeleteDashboard returned error: %v", err)
	}
}

func TestService_DashboardsMutateError(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sendGraphQLResponse(t, w, createDashboardResponse{
			CreateDashboard: createDashboardCreateDashboardCreateDashboardResponse{
				Success: false,
				Code:    "no-buzzing-the-tower",
				Message: "negative ghost rider the pattern is full",
			},
		})
	})

	_, err := client.DashboardsService().Create(ctx, CreateDashboardInput{})
	if err == nil {
		t.Error("Swo.DashboardsMutateErrors expected an error response")
	}
}

// TestService_DashboardsValidationError verifies that a failed create mutation returns a
// structured *DashboardMutationError and that its Error() string surfaces the server code,
// message, and every field-level validation error (including numeric constraints). The
// Terraform provider formats this error as a string, so dropping any of these details from
// the message would silently hide validation feedback from users.
func TestService_DashboardsValidationError(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sendGraphQLResponse(t, w, createDashboardResponse{
			CreateDashboard: createDashboardCreateDashboardCreateDashboardResponse{
				Success: false,
				Code:    "VALIDATION_FAILED",
				Message: "dashboard validation failed",
				ValidationErrors: []createDashboardCreateDashboardCreateDashboardResponseValidationErrorsDashboardValidationError{
					{
						DashboardValidationError: DashboardValidationError{
							Field:   "name",
							Code:    DashboardValidationErrorCodeRequired,
							Message: "name is required",
						},
					},
					{
						DashboardValidationError: DashboardValidationError{
							Field:   "layout",
							Code:    DashboardValidationErrorCodeInvalidRange,
							Message: "layout out of range",
							Constraints: &DashboardValidationErrorConstraintsNumericRangeConstraint{
								Min: Ptr(2.0),
								Max: Ptr(10.0),
							},
						},
					},
				},
			},
		})
	})

	_, err := client.DashboardsService().Create(ctx, CreateDashboardInput{})
	if err == nil {
		t.Fatal("Swo.DashboardsValidationError expected an error response")
	}

	var mutationErr *DashboardMutationError
	if !errors.As(err, &mutationErr) {
		t.Fatalf("Swo.DashboardsValidationError expected *DashboardMutationError, got %T", err)
	}

	if mutationErr.Code != "VALIDATION_FAILED" {
		t.Errorf("Code got = %q, want = %q", mutationErr.Code, "VALIDATION_FAILED")
	}
	if len(mutationErr.ValidationErrors) != 2 {
		t.Fatalf("ValidationErrors length got = %d, want = %d", len(mutationErr.ValidationErrors), 2)
	}

	msg := err.Error()
	for _, want := range []string{
		"code: VALIDATION_FAILED",
		"message: dashboard validation failed",
		"field=name code=REQUIRED message=name is required",
		"field=layout code=INVALID_RANGE message=layout out of range (min=2, max=10)",
	} {
		if !strings.Contains(msg, want) {
			t.Errorf("Error() = %q, want it to contain %q", msg, want)
		}
	}
}

func TestService_DashboardsServerErrors(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", httpErrorResponse)

	_, err := client.DashboardsService().Create(ctx, CreateDashboardInput{})
	if err == nil {
		t.Error("Swo.DashboardsServerErrors expected an error response")
	}
	_, err = client.DashboardsService().Read(ctx, "123")
	if err == nil {
		t.Error("Swo.DashboardsServerErrors expected an error response")
	}
	_, err = client.DashboardsService().Update(ctx, UpdateDashboardInput{})
	if err == nil {
		t.Error("Swo.DashboardsServerErrors expected an error response")
	}
	err = client.DashboardsService().Delete(ctx, "123")
	if err == nil {
		t.Error("Swo.DashboardsServerErrors expected an error response")
	}
}

func TestDashboard_Marshal(t *testing.T) {
	testJSONMarshal(t, &ReadDashboardResult{}, "{}")

	id := uuid.NewString()
	isPrivate := false
	var props any = struct{}{}

	got := ReadDashboardResult{
		Id:        id,
		Name:      dashboardsMockData.fieldName,
		IsPrivate: &isPrivate,
		CreatedAt: dashboardsMockData.fieldUpdatedAt,
		UpdatedAt: dashboardsMockData.fieldUpdatedAt,
		OwnerId:   &dashboardsMockData.fieldOwnerId,
		Layout: []getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardLayout{
			{Id: "123", X: 0, Y: 0, Height: 2, Width: 2},
		},
		Widgets: []getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardWidgetsWidget{
			{Id: "123", Type: "Proportional", Properties: &props},
		},
	}

	want := fmt.Sprintf(`
	{
		"id": "%s",
		"name": "%s",
		"isPrivate": false,
		"createdAt": "%s",
		"updatedAt": "%s",
		"ownerId": "%s",
		"layout": [
			{ "id": "123", "x": 0, "y": 0, "height": 2, "width": 2 }
		],
		"widgets": [
			{ "id": "123", "type": "Proportional", "properties": {} }
		]
	}`,
		id,
		dashboardsMockData.fieldName,
		dashboardsMockData.fieldUpdatedAt.Format(time.RFC3339),
		dashboardsMockData.fieldUpdatedAt.Format(time.RFC3339),
		dashboardsMockData.fieldOwnerId)

	testJSONMarshal(t, got, want)
}
