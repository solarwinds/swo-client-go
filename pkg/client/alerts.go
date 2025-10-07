package client

import (
	"context"
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

// Create creates a new alert with the given definition.
func (as *AlertsService) Create(ctx context.Context, input AlertDefinitionInput) (*CreateAlertDefinitionResult, error) {
	log.Printf("Create alert request. Name: %s", input.Name)

	resp, err := createAlertDefinitionMutation(ctx, as.client.gql, input)
	if err != nil {
		return nil, err
	}

	alertDef := resp.AlertMutations.CreateAlertDefinition
	if alertDef == nil {
		return nil, ErrUnknown
	}

	log.Printf("Create alert success. Id: %s", alertDef.Id)
	return alertDef, nil
}

// Read returns the alert identified by the given Id.
func (as *AlertsService) Read(ctx context.Context, id string) (*ReadAlertDefinitionResult, error) {
	log.Printf("Read alert request. Id: %s", id)

	resp, err := getAlertDefinitionById(ctx, as.client.gql, id)
	if err != nil {
		return nil, err
	}

	alertDefs := resp.AlertQueries.AlertDefinitions.AlertDefinitions
	if len(alertDefs) == 0 {
		return nil, ErrNotFound
	}

	log.Printf("Read alert success. Id: %s", id)
	return &alertDefs[0], nil
}

// Update updates the alert with the given id.
func (as *AlertsService) Update(ctx context.Context, id string, input AlertDefinitionInput) (*UpdateAlertDefinitionResult, error) {
	log.Printf("Update alert request. Id: %s", id)

	resp, err := updateAlertDefinitionMutation(ctx, as.client.gql, input, id)
	if err != nil {
		return nil, err
	}

	result := resp.AlertMutations.UpdateAlertDefinition
	if result == nil {
		log.Printf("Alert not found. Id: %s", id)
		return nil, ErrNotFound
	}

	log.Printf("Update alert success. Id: %s", id)
	return result, nil
}

// Delete deletes the alert with the given id.
func (as *AlertsService) Delete(ctx context.Context, id string) error {
	log.Printf("Delete alert request. Id: %s", id)

	r, err := deleteAlertDefinitionMutation(ctx, as.client.gql, id)
	if err != nil {
		return err
	}

	idPtr := r.AlertMutations.DeleteAlertDefinition
	if idPtr == nil || *idPtr != id {
		log.Printf("Alert not found. Id: %s", id)
		return ErrNotFound
	}

	log.Printf("Delete alert success. Id: %s", id)
	return nil
}
