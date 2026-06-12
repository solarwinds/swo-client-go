package client

import (
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestSwoService_ReadCircleCIConnection(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	inputId := uuid.NewString()
	output := &getCircleCIConnectionResponse{
		Vcs: getCircleCIConnectionVcsVcsQueries{
			GetCircleCIConnection: &getCircleCIConnectionVcsVcsQueriesGetCircleCIConnectionVcsCircleCIConnection{
				Id:          inputId,
				Name:        "test-connection",
				SecretToken: "secret-abc",
			},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__getCircleCIConnectionInput](r)
		if err != nil {
			t.Errorf("Swo.ReadCircleCIConnection error: %v", err)
		}

		got := gqlInput.Id
		want := inputId

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, output)
	})

	got, err := client.CircleCIIntegrationService().Read(ctx, inputId)
	if err != nil {
		t.Errorf("Swo.ReadCircleCIConnection returned error: %v", err)
	}

	want := output.Vcs.GetCircleCIConnection

	if !testObjects(t, *got, *want) {
		t.Errorf("Swo.ReadCircleCIConnection returned %+v, wanted %+v", got, want)
	}
}

func TestSwoService_CreateCircleCIConnection(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	inputName := "my-circleci"
	inputToken := "cci-token-123"
	output := &createCircleCIConnectionResponse{
		Vcs: createCircleCIConnectionVcsVcsMutations{
			CreateCircleCIConnection: createCircleCIConnectionVcsVcsMutationsCreateCircleCIConnectionVcsCircleCIConnection{
				Id:          uuid.NewString(),
				Name:        inputName,
				SecretToken: "secret-xyz",
			},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createCircleCIConnectionInput](r)
		if err != nil {
			t.Errorf("Swo.CreateCircleCIConnection error: %v", err)
		}

		if gqlInput.Name != inputName {
			t.Errorf("Request name got = %+v, want = %+v", gqlInput.Name, inputName)
		}
		if *gqlInput.ApiToken != inputToken {
			t.Errorf("Request apiToken got = %+v, want = %+v", *gqlInput.ApiToken, inputToken)
		}

		sendGraphQLResponse(t, w, output)
	})

	got, err := client.CircleCIIntegrationService().Create(ctx, inputName, &inputToken)
	if err != nil {
		t.Errorf("Swo.CreateCircleCIConnection returned error: %v", err)
	}

	want := output.Vcs.CreateCircleCIConnection

	if !testObjects(t, *got, want) {
		t.Errorf("Swo.CreateCircleCIConnection returned %+v, want %+v", got, want)
	}
}

func TestSwoService_UpdateCircleCIConnection(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	inputId := uuid.NewString()
	inputName := "updated-name"
	inputToken := "new-token"
	output := &updateCircleCIConnectionResponse{
		Vcs: updateCircleCIConnectionVcsVcsMutations{
			UpdateCircleCIConnection: updateCircleCIConnectionVcsVcsMutationsUpdateCircleCIConnectionVcsCircleCIConnection{
				Id:          inputId,
				Name:        inputName,
				SecretToken: "secret-xyz",
			},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__updateCircleCIConnectionInput](r)
		if err != nil {
			t.Errorf("Swo.UpdateCircleCIConnection error: %v", err)
		}

		if gqlInput.Id != inputId {
			t.Errorf("Request id got = %+v, want = %+v", gqlInput.Id, inputId)
		}
		if *gqlInput.Name != inputName {
			t.Errorf("Request name got = %+v, want = %+v", *gqlInput.Name, inputName)
		}
		if *gqlInput.ApiToken != inputToken {
			t.Errorf("Request apiToken got = %+v, want = %+v", *gqlInput.ApiToken, inputToken)
		}

		sendGraphQLResponse(t, w, output)
	})

	got, err := client.CircleCIIntegrationService().Update(ctx, inputId, &inputName, &inputToken)
	if err != nil {
		t.Errorf("Swo.UpdateCircleCIConnection returned error: %v", err)
	}

	want := output.Vcs.UpdateCircleCIConnection

	if !testObjects(t, *got, want) {
		t.Errorf("Swo.UpdateCircleCIConnection returned %+v, want %+v", got, want)
	}
}

func TestSwoService_DeleteCircleCIConnection(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	inputId := uuid.NewString()
	output := &deleteCircleCIConnectionResponse{
		Vcs: deleteCircleCIConnectionVcsVcsMutations{
			DeleteCircleCIConnection: deleteCircleCIConnectionVcsVcsMutationsDeleteCircleCIConnectionVcsDeleteResponse{
				Success: true,
			},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__deleteCircleCIConnectionInput](r)
		if err != nil {
			t.Errorf("Swo.DeleteCircleCIConnection error: %v", err)
		}

		got := gqlInput.Id
		want := inputId

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, output)
	})

	if err := client.CircleCIIntegrationService().Delete(ctx, inputId); err != nil {
		t.Errorf("Swo.DeleteCircleCIConnection returned error: %v", err)
	}
}
