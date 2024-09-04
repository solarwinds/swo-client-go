package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestSwoService_ReadLogFilter(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := "123"

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__getLogFilterByIdInput](r)
		if err != nil {
			t.Errorf("Swo.ReadLogFilter error: %v", err)
		}

		got := gqlInput.Input.Id
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, getLogFilterByIdResponse{
			GetExclusionFilter: &ReadLogFilterResult{
				Name:        "swo-client-go - logFilter",
				Description: Ptr("logFilter description"),
				Expressions: []getLogFilterByIdGetExclusionFilterExpressionsExclusionFilterExpression{
					{
						Kind:       ExclusionFilterExpressionKindString,
						Expression: "test string",
					},
				},
			},
		})
	})

	got, err := client.LogFilterService().Read(ctx, "123")
	if err != nil {
		t.Errorf("Swo.ReadLogFilter returned error: %v", err)
	}

	want := &ReadLogFilterResult{
		Name:        "swo-client-go - logFilter",
		Description: Ptr("logFilter description"),
		Expressions: []getLogFilterByIdGetExclusionFilterExpressionsExclusionFilterExpression{
			{
				Kind:       ExclusionFilterExpressionKindString,
				Expression: "test string",
			},
		},
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.ReadLogFilter returned %+v, wanted %+v", got, want)
	}
}

func TestSwoService_CreateLogFilter(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := &CreateExclusionFilterInput{
		Name:        "swo-client-go - logFilter",
		Description: "logFilter description",
		Expressions: []CreateExclusionFilterExpressionInput{
			{
				Kind:       ExclusionFilterExpressionKindString,
				Expression: "test string",
			},
		},
	}

	id := uuid.NewString()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createLogFilterInput](r)
		if err != nil {
			t.Errorf("Swo.CreateLogFilter error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, *want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, createLogFilterResponse{
			CreateExclusionFilter: createLogFilterCreateExclusionFilterCreateExclusionFilterResponse{
				Code:    ExclusionFilterResponseCodeOk,
				Success: true,
				Message: "ok",
				ExclusionFilter: &createLogFilterCreateExclusionFilterCreateExclusionFilterResponseExclusionFilter{
					Id:          id,
					Name:        input.Name,
					Description: &input.Description,
					Expressions: []createLogFilterCreateExclusionFilterCreateExclusionFilterResponseExclusionFilterExpressionsExclusionFilterExpression{
						{
							Kind:       ExclusionFilterExpressionKindString,
							Expression: "test string",
						},
					},
				},
			},
		})
	})

	got, err := client.LogFilterService().Create(ctx, *input)
	if err != nil {
		t.Errorf("Swo.CreateLogFilter returned error: %v", err)
	}

	want := &CreateLogFilterResult{
		Id:          id,
		Name:        input.Name,
		Description: &input.Description,
		Expressions: []createLogFilterCreateExclusionFilterCreateExclusionFilterResponseExclusionFilterExpressionsExclusionFilterExpression{
			{
				Kind:       ExclusionFilterExpressionKindString,
				Expression: "test string",
			},
		},
	}

	if !testObjects(t, got, want) {
		t.Errorf("Swo.CreateLogFilter returned %+v, want %+v", got, want)
	}
}

func TestSwoService_UpdateLogFilter(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := UpdateExclusionFilterInput{
		Id:          "123",
		Name:        "swo-client-go - logFilter",
		Description: "logFilter description",
		Expressions: []UpdateExclusionFilterExpressionInput{
			{
				Kind:       ExclusionFilterExpressionKindString,
				Expression: "test string",
			},
		},
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__updateLogFilterInput](r)
		if err != nil {
			t.Errorf("Swo.UpdateLogFilter error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, updateLogFilterUpdateExclusionFilterGenericExclusionFilterMutationResponse{
			Code:    "OK",
			Success: true,
			Message: "ok",
		})
	})

	if err := client.LogFilterService().Update(ctx, input); err != nil {
		t.Errorf("Swo.UpdateLogFilter returned error: %v", err)
	}
}

func TestSwoService_DeleteLogFilter(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := "123"

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__deleteLogFilterInput](r)
		if err != nil {
			t.Errorf("Swo.DeleteLogFilter error: %v", err)
		}

		got := gqlInput.Input.Id
		want := input

		if !testObjects(t, got, want) {
			t.Errorf("Swo.DeleteLogFilter: Request got = %+v, want %+v", got, want)
		}

		sendGraphQLResponse(t, w, deleteLogFilterDeleteExclusionFilterGenericExclusionFilterMutationResponse{
			Code:    "OK",
			Success: true,
			Message: "ok",
		})
	})

	err := client.LogFilterService().Delete(ctx, input)
	if err != nil {
		t.Errorf("Swo.DeleteLogFilter returned error: %v", err)
	}
}

func TestSwoService_LogFilterServerErrors(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", httpErrorResponse)

	_, err := client.LogFilterService().Create(ctx, CreateExclusionFilterInput{})
	if err == nil {
		t.Error("Swo.LogFilterServerErrors expected an error response")
	}
	_, err = client.LogFilterService().Read(ctx, "123")
	if err == nil {
		t.Error("Swo.LogFilterServerErrors expected an error response")
	}
	err = client.LogFilterService().Update(ctx, UpdateExclusionFilterInput{})
	if err == nil {
		t.Error("Swo.LogFilterServerErrors expected an error response")
	}
	err = client.LogFilterService().Delete(ctx, "123")
	if err == nil {
		t.Error("Swo.LogFilterServerErrors expected an error response")
	}
}

func TestSwoService_LogFilterDupicateEntryError(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := &CreateExclusionFilterInput{
		Name:        "swo-client-go - logFilter",
		Description: "logFilter description",
		Expressions: []CreateExclusionFilterExpressionInput{
			{
				Kind:       ExclusionFilterExpressionKindString,
				Expression: "test string",
			},
		},
	}

	message := "Global filter for this org already exists"

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createLogFilterInput](r)
		if err != nil {
			t.Errorf("Swo.LogFilterDupicateEntryError error: %v", err)
		}

		got := gqlInput.Input
		want := input

		if !testObjects(t, got, *want) {
			t.Errorf("Request got = %+v, want = %+v", got, want)
		}

		sendGraphQLResponse(t, w, createLogFilterResponse{
			CreateExclusionFilter: createLogFilterCreateExclusionFilterCreateExclusionFilterResponse{
				Code:            ExclusionFilterResponseCodeDuplicateEntry,
				Success:         false,
				Message:         message,
				ExclusionFilter: nil,
			},
		})
	})

	_, err := client.LogFilterService().Create(ctx, *input)

	if err == nil {
		t.Error("Swo.LogFilterDupicateEntryError expected an error response")
	}

	want := fmt.Errorf("create LogFilter failed. code=%s message=%s", ExclusionFilterResponseCodeDuplicateEntry, message)

	if !testObjects(t, err.Error(), want.Error()) {
		t.Errorf("Swo.LogFilterDupicateEntryError returned %+v, want %+v", err, want)
	}
}
