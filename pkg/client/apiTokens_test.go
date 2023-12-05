package client

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestSwoService_ReadApiToken(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	inputId := uuid.NewString()
	output := &getApiTokenByIdResponse{
		User: getApiTokenByIdUserAuthenticatedUser{
			CurrentOrganization: getApiTokenByIdUserAuthenticatedUserCurrentOrganization{
				Tokens: []getApiTokenByIdUserAuthenticatedUserCurrentOrganizationTokensToken{
					{
						Id:              inputId,
						Name:            Ptr("swo-client-go - apiToken"),
						Token:           Ptr("123"),
						ObfuscatedToken: Ptr("123"),
						AccessLevel:     Ptr(TokenAccessLevelFull),
						Attributes: []getApiTokenByIdUserAuthenticatedUserCurrentOrganizationTokensTokenAttributesTokenAttribute{
							{
								Key:   "test-key",
								Value: "test-value",
							},
						},
						Enabled:       Ptr(true),
						Type:          Ptr("public-api"),
						Secure:        Ptr(true),
						CreatedBy:     Ptr("123"),
						CreatedByName: Ptr("user name"),
						CreatedAt:     Ptr(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
						UpdatedAt:     Ptr(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__getApiTokenByIdInput](r)
		if err != nil {
			t.Errorf("Swo.ReadApiToken error: %v", err)
		}

		got := gqlInput.Id
		want := inputId

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, output)
	})

	got, err := client.ApiTokenService().Read(ctx, inputId)
	if err != nil {
		t.Errorf("Swo.ReadApiToken returned error: %v", err)
	}

	want := output.User.CurrentOrganization.Tokens[0]

	if !testObjects(t, *got, want) {
		t.Errorf("Swo.ReadApiToken returned %+v, wanted %+v", got, want)
	}
}

func TestSwoService_CreateApiToken(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := &CreateTokenInput{
		Name:        "swo-client-go - apiToken",
		AccessLevel: TokenAccessLevelFull,
		Type:        Ptr("public-api"),
		Attributes: []TokenAttributeInput{
			{
				Key:   "test-key",
				Value: "test-value",
			},
		},
	}
	output := &createTokenMutationResponse{
		CreateToken: &createTokenMutationCreateTokenCreateTokenResponse{
			Success: true,
			Token: &createTokenMutationCreateTokenCreateTokenResponseToken{
				Id:          uuid.NewString(),
				Name:        Ptr(input.Name),
				Token:       Ptr("123"),
				AccessLevel: Ptr(TokenAccessLevelFull),
				Type:        Ptr("public-api"),
				CreatedAt:   Ptr(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createTokenMutationInput](r)
		if err != nil {
			t.Errorf("Swo.CreateApiToken error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, *want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, output)
	})

	got, err := client.ApiTokenService().Create(ctx, *input)
	if err != nil {
		t.Errorf("Swo.CreateApiToken returned error: %v", err)
	}

	want := output.CreateToken.Token

	if !testObjects(t, got, want) {
		t.Errorf("Swo.CreateApiToken returned %+v, want %+v", got, want)
	}
}

func TestSwoService_UpdateApiToken(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := UpdateTokenInput{
		Id:          uuid.NewString(),
		Name:        Ptr("swo-client-go - apiToken"),
		AccessLevel: Ptr(TokenAccessLevelFull),
		Type:        Ptr("public-api"),
		Attributes: []TokenAttributeInput{
			{
				Key:   "test-key",
				Value: "test-value",
			},
		},
	}
	output := &updateTokenMutationResponse{
		UpdateToken: &updateTokenMutationUpdateTokenUpdateTokenResponse{
			Success: true,
			Token: &updateTokenMutationUpdateTokenUpdateTokenResponseToken{
				Id:              input.Id,
				Name:            input.Name,
				ObfuscatedToken: Ptr("123"),
				AccessLevel:     input.AccessLevel,
				Type:            input.Type,
				Enabled:         Ptr(true),
				UpdatedAt:       Ptr(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__updateTokenMutationInput](r)
		if err != nil {
			t.Errorf("Swo.UpdateApiToken error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, output)
	})

	if err := client.ApiTokenService().Update(ctx, input); err != nil {
		t.Errorf("Swo.UpdateApiToken returned error: %v", err)
	}
}

func TestSwoService_DeleteApiToken(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := uuid.NewString()
	output := &deleteTokenMutationResponse{
		DeleteToken: &deleteTokenMutationDeleteTokenDeleteTokenResponse{
			Success: true,
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__deleteTokenMutationInput](r)
		if err != nil {
			t.Errorf("Swo.DeleteApiToken error: %v", err)
		}

		got := gqlInput.Input.Id
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Swo.DeleteApiToken: Request got = %+v, want %+v", got, want)
		}

		sendGraphQLResponse(t, w, output)
	})

	if err := client.ApiTokenService().Delete(ctx, input); err != nil {
		t.Errorf("Swo.DeleteApiToken returned error: %v", err)
	}
}

func TestSwoService_ApiTokenServerErrors(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", httpErrorResponse)

	if _, err := client.ApiTokenService().Create(ctx, CreateTokenInput{}); err == nil {
		t.Error("Swo.ApiTokenServerErrors expected an error response")
	}
	if _, err := client.ApiTokenService().Read(ctx, "123"); err == nil {
		t.Error("Swo.ApiTokenServerErrors expected an error response")
	}
	if err := client.ApiTokenService().Update(ctx, UpdateTokenInput{}); err == nil {
		t.Error("Swo.ApiTokenServerErrors expected an error response")
	}
	if err := client.ApiTokenService().Delete(ctx, "123"); err == nil {
		t.Error("Swo.ApiTokenServerErrors expected an error response")
	}
}
