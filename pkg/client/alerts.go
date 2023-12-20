package client

import (
	"context"
	"fmt"
	"log"
)

type AlertsService service

// Exported Create types
type CreateAlertDefinitionResult = createAlertDefinitionMutationAlertMutationsCreateAlertDefinition

// Exported Update types
type UpdateAlertDefinitionResult = updateAlertDefinitionMutationAlertMutationsUpdateAlertDefinition

// Exported Read types
type ReadAlertDefinitionResult = getAlertDefinitionByIdAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinition
type ReadAlertConditionResult = getAlertDefinitionByIdAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpression
type ReadAlertActionResult = getAlertDefinitionByIdAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionActionsAlertAction
type ReadAlertMuteInfoResult = getAlertDefinitionByIdAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionMuteInfo
type ReadAlertUserResult = getAlertDefinitionByIdAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionUser
type ReadAlertConditionLinkResult = getAlertDefinitionByIdAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionLinksNamedLinks
type ReadAlertConditionValueResult = getAlertDefinitionByIdAlertQueriesAlertDefinitionsAlertDefinitionsResultAlertDefinitionsAlertDefinitionFlatConditionFlatAlertConditionExpressionValueFlatAlertConditionNode

type AlertsCommunicator interface {
	Create(context.Context, AlertDefinitionInput) (*CreateAlertDefinitionResult, error)
	Read(context.Context, string) (*ReadAlertDefinitionResult, error)
	Update(context.Context, string, AlertDefinitionInput) (*UpdateAlertDefinitionResult, error)
	Delete(context.Context, string) error
}

func newAlertsService(c *Client) *AlertsService {
	return &AlertsService{c}
}

// Creates a new alert with the given definition.
func (as *AlertsService) Create(ctx context.Context, input AlertDefinitionInput) (*CreateAlertDefinitionResult, error) {
	log.Printf("Create alert request. Name: %s", input.Name)

	resp, err := createAlertDefinitionMutation(ctx, as.client.gql, input)

	if err != nil {
		return nil, err
	}

	alertDef := resp.AlertMutations.CreateAlertDefinition
	log.Printf("Create alert success. Id: %s", alertDef.Id)

	return alertDef, nil
}

// Returns the alert identified by the given Id.
func (as *AlertsService) Read(ctx context.Context, id string) (*ReadAlertDefinitionResult, error) {
	log.Printf("Read alert request. Id: %s", id)

	resp, err := getAlertDefinitionById(ctx, as.client.gql, id)

	if err != nil {
		return nil, err
	}

	alertDefs := resp.AlertQueries.AlertDefinitions.AlertDefinitions

	if len(alertDefs) == 0 {
		return nil, fmt.Errorf("alert not found. id: %s", id)
	}

	alertDef := resp.AlertQueries.AlertDefinitions.AlertDefinitions[0]
	log.Printf("Read alert success. Id: %s", id)

	return &alertDef, nil
}

// Updates the alert with the given id.
func (as *AlertsService) Update(ctx context.Context, id string, input AlertDefinitionInput) (*UpdateAlertDefinitionResult, error) {
	log.Printf("Update alert request. Id: %s", id)

	resp, err := updateAlertDefinitionMutation(ctx, as.client.gql, input, id)
	if err != nil {
		return nil, err
	}

	log.Printf("Update alert success. Id: %s", id)

	return resp.AlertMutations.UpdateAlertDefinition, nil
}

// Deletes the alert with the given id.
func (as *AlertsService) Delete(ctx context.Context, id string) error {
	log.Printf("Delete alert request. Id: %s", id)

	_, err := deleteAlertDefinitionMutation(ctx, as.client.gql, id)

	if err != nil {
		return err
	}

	log.Printf("Delete alert success. Id: %s", id)

	return nil
}
