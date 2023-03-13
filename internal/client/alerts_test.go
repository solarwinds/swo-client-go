package client

import (
	"net/http"
	"testing"
)

var (
	mockAlertId              = "43a1743c-91ca-43ee-a37e-df01902d2dc4"
	mockAlertName            = "swo-client-go_test-alert"
	mockAlertDescription     = "testing alert creation from swo-client-go"
	mockAlertDefinitionInput = func(name string, description string) AlertDefinitionInput {
		operatorGt := ">"
		operatorMax := "MAX"

		alertExpression := AlertFilterExpressionInput{Operation: FilterOperationEq}
		entityFilter := AlertConditionNodeEntityFilterInput{Type: "DeviceVolume"}

		dataTypeString := "string"
		dataTypeNumber := "number"

		oneDValue := "1d"
		Ninety := "90"

		fieldName := "Orion.NPM.InterfaceTraffic.InTotalBytes"

		return AlertDefinitionInput{
			Name:        name,
			Description: &description,
			Enabled:     true,
			Severity:    AlertSeverityInfo,
			Condition: []AlertConditionNodeInput{
				{Id: 0, Type: "binaryOperator", Operator: &operatorGt, OperandIds: []int{1, 4}},
				{Id: 1, Type: "aggregationOperator", Operator: &operatorMax, OperandIds: []int{2, 3}, MetricFilter: &alertExpression},
				{Id: 2, Type: "metricField", EntityFilter: &entityFilter, FieldName: &fieldName},
				{Id: 3, Type: "constantValue", DataType: &dataTypeString, Value: &oneDValue},
				{Id: 4, Type: "constantValue", DataType: &dataTypeNumber, Value: &Ninety},
			},
		}
	}
	mockCreateAlertDefinitionResult = func(id string, name string, description string) *CreateAlertDefinitionResult {
		return &CreateAlertDefinitionResult{
			Id:             id,
			Name:           name,
			Description:    &description,
			Enabled:        true,
			OrganizationId: "140638900734749696",
			Severity:       "INFO",
			Actions:        []createAlertDefinitionAlertMutationsCreateAlertDefinitionActionsAlertAction{},
			FlatCondition: []createAlertDefinitionAlertMutationsCreateAlertDefinitionFlatConditionFlatAlertConditionExpression{
				{
					Id: "935b93f6-f94f-4b25-98a6-e66bbf80eaee",
				},
				{
					Id: "0f9212ff-c437-4496-aabd-72a3d0c4dea0",
				},
				{
					Id: "9b3f1343-4936-40d2-bb2c-7ee8a7612f46",
				},
				{
					Id: "8fc84a11-dece-4561-9b36-573c4e38929f",
				},
				{
					Id: "fa655d66-673c-444c-b368-4c173585699d",
				},
			},
		}
	}
	mockReadAlertDefinitionResult = func(id string, description string) *getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinition {
		operatorGT := ">"
		operatorMAX := "MAX"
		metricFieldName := "Orion.NPM.InterfaceTraffic.InTotalBytes"

		return &getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinition{
			Id:                id,
			Name:              "terraform-provider-swo-test-alert",
			Description:       &description,
			Enabled:           false,
			OrganizationId:    "140638900734749696",
			UserId:            "151686710111094784",
			Severity:          "INFO",
			Triggered:         false,
			TargetEntityTypes: []string{"DeviceVolume"},
			MuteInfo: getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionMuteInfo{
				Muted: false,
			},
			Actions:             []getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionActionsAlertAction{},
			TriggerResetActions: false,
			ConditionType:       "ENTITY_METRIC",
			FlatCondition: []getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpression{
				{
					Id: "935b93f6-f94f-4b25-98a6-e66bbf80eaee",
					Links: []getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionLinksNamedLinks{
						{
							Name: "operands",
							Values: []string{
								"0f9212ff-c437-4496-aabd-72a3d0c4dea0",
								"9b3f1343-4936-40d2-bb2c-7ee8a7612f46",
							},
						},
					},
					Value: &getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionValueFlatAlertConditionNode{
						FieldName: nil,
						Operator:  &operatorGT,
						Type:      "binaryOperator",
						Query:     nil,
					},
				},
				{
					Id: "0f9212ff-c437-4496-aabd-72a3d0c4dea0",
					Links: []getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionLinksNamedLinks{
						{
							Name: "operands",
							Values: []string{
								"8fc84a11-dece-4561-9b36-573c4e38929f",
								"fa655d66-673c-444c-b368-4c173585699d",
							},
						},
					},
					Value: &getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionValueFlatAlertConditionNode{
						FieldName: nil,
						Operator:  &operatorMAX,
						Type:      "aggregationOperator",
						Query:     nil,
					},
				},
				{
					Id:    "9b3f1343-4936-40d2-bb2c-7ee8a7612f46",
					Links: []getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionLinksNamedLinks{},
					Value: &getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionValueFlatAlertConditionNode{
						FieldName: nil,
						Operator:  nil,
						Type:      "constantValue",
						Query:     nil,
					},
				},
				{
					Id:    "8fc84a11-dece-4561-9b36-573c4e38929f",
					Links: []getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionLinksNamedLinks{},
					Value: &getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionValueFlatAlertConditionNode{
						FieldName: &metricFieldName,
						Operator:  nil,
						Type:      "metricField",
						Query:     nil,
					},
				},
				{
					Id:    "fa655d66-673c-444c-b368-4c173585699d",
					Links: []getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionLinksNamedLinks{},
					Value: &getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionValueFlatAlertConditionNode{
						FieldName: nil,
						Operator:  nil,
						Type:      "constantValue",
						Query:     nil,
					},
				},
			},
		}
	}
	mockUpdateAlertDefinitionResult = func(id string, name string, description string) *UpdateAlertDefinitionResult {
		return &UpdateAlertDefinitionResult{
			Id:             id,
			Name:           name,
			Description:    &description,
			Enabled:        true,
			OrganizationId: "140638900734749696",
			Severity:       "INFO",
			Actions:        []updateAlertDefinitionAlertMutationsUpdateAlertDefinitionActionsAlertAction{},
			FlatCondition: []updateAlertDefinitionAlertMutationsUpdateAlertDefinitionFlatConditionFlatAlertConditionExpression{
				{
					Id: "935b93f6-f94f-4b25-98a6-e66bbf80eaee",
				},
				{
					Id: "0f9212ff-c437-4496-aabd-72a3d0c4dea0",
				},
				{
					Id: "9b3f1343-4936-40d2-bb2c-7ee8a7612f46",
				},
				{
					Id: "8fc84a11-dece-4561-9b36-573c4e38929f",
				},
				{
					Id: "fa655d66-673c-444c-b368-4c173585699d",
				},
			},
		}
	}
)

func TestCreateAlert(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__createAlertDefinitionInput](r)
		if err != nil {
			t.Errorf("Swo.CreateAlert returned error: %v", err)
		}

		sendGraphQLResponse(t, w, createAlertDefinitionResponse{
			AlertMutations: createAlertDefinitionAlertMutations{
				CreateAlertDefinition: mockCreateAlertDefinitionResult(mockAlertId, gqlInput.Definition.Name, *gqlInput.Definition.Description),
			},
		})
	})

	got, err := client.AlertsService().Create(ctx, mockAlertDefinitionInput(mockAlertName, mockAlertDescription))
	if err != nil {
		t.Errorf("Swo.ReadAlert error: %v", err)
		return
	}

	want := mockCreateAlertDefinitionResult(mockAlertId, mockAlertName, mockAlertDescription)

	if !testObjects(t, got, want) {
		t.Errorf("Swo.ReadAlert returned %+v, wanted %+v", got, want)
	}
}

func TestReadAlert(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__getAlertDefinitionsInput](r)
		if err != nil {
			t.Errorf("Swo.ReadAlert returned error: %v", err)
		}

		sendGraphQLResponse(t, w, getAlertDefinitionsResponse{
			AlertQueries: getAlertDefinitionsAlertQueries{
				AlertDefinitions: getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResult{
					AlertDefinitions: []getAlertDefinitionsAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinition{
						*mockReadAlertDefinitionResult(*gqlInput.Filter.Id, mockAlertDescription),
					},
				},
			},
		})
	})

	got, err := client.AlertsService().Read(ctx, mockAlertId)
	if err != nil {
		t.Errorf("Swo.ReadAlert error: %v", err)
		return
	}

	want := mockReadAlertDefinitionResult(mockAlertId, mockAlertDescription)

	if !testObjects(t, got, want) {
		t.Errorf("Swo.ReadAlert returned %+v, wanted %+v", got, want)
	}
}

func TestUpdateAlert(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := mockAlertDefinitionInput(mockAlertName, mockAlertDescription)

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__updateAlertDefinitionInput](r)
		if err != nil {
			t.Errorf("Swo.ReadAlert returned error: %v", err)
		}

		got := gqlInput.Definition
		want := input

		if !testObjects(t, want, got) {
			t.Errorf("Request input = %+v, want %+v", got, want)
		}

		sendGraphQLResponse(t, w, updateAlertDefinitionAlertMutations{
			UpdateAlertDefinition: mockUpdateAlertDefinitionResult(gqlInput.UpdateAlertDefinitionId, got.Name, *got.Description),
		})
	})

	_, err := client.AlertsService().Update(ctx, mockAlertId, input)
	if err != nil {
		t.Errorf("Swo.UpdateNotification returned error: %v", err)
	}
}

func TestDeleteAlert(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	input := deleteAlertDefinitionAlertMutations{
		DeleteAlertDefinition: &mockAlertId,
	}

	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		gqlInput, err := getGraphQLInput[__deleteAlertDefinitionInput](r)
		if err != nil {
			t.Errorf("Swo.DeleteAlert returned error: %v", err)
		}

		got := gqlInput.DeleteAlertDefinitionId
		want := *input.DeleteAlertDefinition

		if !testObjects(t, got, want) {
			t.Errorf("Swo.DeleteAlert: got = %s, want %s", got, want)
		}

		sendGraphQLResponse(t, w, deleteAlertDefinitionResponse{
			AlertMutations: deleteAlertDefinitionAlertMutations{
				DeleteAlertDefinition: &mockAlertId,
			},
		})
	})

	err := client.AlertsService().Delete(ctx, *input.DeleteAlertDefinition)
	if err != nil {
		t.Errorf("Swo.DeleteAlert returned error: %v", err)
	}
}

func TestSwoService_AlertsServerErrors(t *testing.T) {
	ctx, client, server, _, teardown := setup()
	defer teardown()

	server.HandleFunc("/", httpErrorResponse)

	_, err := client.AlertsService().Create(ctx, AlertDefinitionInput{})
	if err == nil {
		t.Error("Swo.AlertServerErrors expected an error response")
	}
	_, err = client.AlertsService().Read(ctx, "123")
	if err == nil {
		t.Error("Swo.AlertServerErrors expected an error response")
	}
	_, err = client.AlertsService().Update(ctx, "123", AlertDefinitionInput{})
	if err == nil {
		t.Error("Swo.AlertServerErrors expected an error response")
	}
	err = client.AlertsService().Delete(ctx, "123")
	if err == nil {
		t.Error("Swo.AlertServerErrors expected an error response")
	}
}
