package client

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestSwoService_ReadUri(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__getUriByIdInput](r)
		if err != nil {
			t.Errorf("Swo.ReadUri error: %v", err)
		}

		got := gqlInput.Id
		want := "123"

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, getUriByIdResponse{
			Entities: getUriByIdEntitiesEntityQueries{},
		})
	})

	got, err := client.UriService().Read(ctx, "123")
	if err != nil {
		t.Errorf("Swo.ReadUri returned error: %v", err)
	}

	want := &getUriByIdResponse{
		Entities: getUriByIdEntitiesEntityQueries{},
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.ReadUri returned %+v, wanted %+v", got, want)
	}
}

func TestSwoService_CreateUri(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input, err := GetObjectFromFile[CreateUriInput]("uri_test.json")
	if err != nil {
		t.Errorf("Swo.CreateUri error: %v", err)
	}

	id := uuid.NewString()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createUriMutationInput](r)
		if err != nil {
			t.Errorf("Swo.CreateUri error: %v", err)
		}

		got := gqlInput.Input
		want := *input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, createUriMutationResponse{
			Dem: createUriMutationDemDemMutations{
				CreateUri: createUriMutationDemDemMutationsCreateUriCreateUriResponse{
					Id: id,
				},
			},
		})
	})

	got, err := client.UriService().Create(ctx, *input)
	if err != nil {
		t.Errorf("Swo.CreateUri returned error: %v", err)
	}

	want := &CreateUriResult{
		Id: id,
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.CreateUri returned %+v, want %+v", got, want)
	}
}

func TestSwoService_UpdateUri(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := UpdateUriInput{
		Id:         "123",
		Name:       "swo-client-go - uri",
		IpOrDomain: "www.solarwinds.com",
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__updateUriMutationInput](r)
		if err != nil {
			t.Errorf("Swo.UpdateUri error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, updateUriMutationResponse{
			Dem: updateUriMutationDemDemMutations{
				UpdateUri: updateUriMutationDemDemMutationsUpdateUriUpdateUriResponse{
					Id: input.Id,
				},
			},
		})
	})

	if err := client.UriService().Update(ctx, input); err != nil {
		t.Errorf("Swo.UpdateUri returned error: %v", err)
	}
}

func TestSwoService_DeleteUri(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := "123"

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__deleteUriMutationInput](r)
		if err != nil {
			t.Errorf("Swo.DeleteUri error: %v", err)
		}

		got := gqlInput.Input.Id
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Swo.DeleteUri: Request got = %+v, want %+v", got, want)
		}

		sendGraphQLResponse(t, w, deleteUriMutationResponse{
			Dem: deleteUriMutationDemDemMutations{
				DeleteUri: deleteUriMutationDemDemMutationsDeleteUriDeleteUriResponse{
					Id: input,
				},
			},
		})
	})

	err := client.UriService().Delete(ctx, input)
	if err != nil {
		t.Errorf("Swo.DeleteUri returned error: %v", err)
	}
}

func TestSwoService_UriServerErrors(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", httpErrorResponse)

	_, err := client.UriService().Create(ctx, CreateUriInput{})
	if err == nil {
		t.Error("Swo.UriServerErrors expected an error response")
	}
	_, err = client.UriService().Read(ctx, "123")
	if err == nil {
		t.Error("Swo.UriServerErrors expected an error response")
	}
	err = client.UriService().Update(ctx, UpdateUriInput{})
	if err == nil {
		t.Error("Swo.UriServerErrors expected an error response")
	}
	err = client.UriService().Delete(ctx, "123")
	if err == nil {
		t.Error("Swo.UriServerErrors expected an error response")
	}
}
