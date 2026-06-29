package client

import (
	"context"
	"fmt"
	"log"
)

type CircleCIIntegrationService service

type CreateCircleCIConnectionResult = createCircleCIConnectionVcsVcsMutationsCreateCircleCIConnectionVcsCircleCIConnection
type ReadCircleCIConnectionResult = getCircleCIConnectionVcsVcsQueriesGetCircleCIConnectionVcsCircleCIConnection
type UpdateCircleCIConnectionResult = updateCircleCIConnectionVcsVcsMutationsUpdateCircleCIConnectionVcsCircleCIConnection

type CircleCIIntegrationCommunicator interface {
	Create(ctx context.Context, name string, apiToken *string) (*CreateCircleCIConnectionResult, error)
	Read(ctx context.Context, id string) (*ReadCircleCIConnectionResult, error)
	Update(ctx context.Context, id string, name *string, apiToken *string) (*UpdateCircleCIConnectionResult, error)
	Delete(ctx context.Context, id string) error
}

func newCircleCIIntegrationService(c *Client) *CircleCIIntegrationService {
	return &CircleCIIntegrationService{c}
}

// Creates a new CircleCI connection with the given name and optional API token.
func (s *CircleCIIntegrationService) Create(ctx context.Context, name string, apiToken *string) (*CreateCircleCIConnectionResult, error) {
	log.Printf("create CircleCI connection request. name=%s", name)

	resp, err := createCircleCIConnection(ctx, s.client.gql, name, apiToken)
	if err != nil {
		return nil, err
	}

	result := &resp.Vcs.CreateCircleCIConnection
	log.Printf("create CircleCI connection success. id=%s", result.Id)

	return result, nil
}

// Returns the CircleCI connection with the given id.
func (s *CircleCIIntegrationService) Read(ctx context.Context, id string) (*ReadCircleCIConnectionResult, error) {
	log.Printf("read CircleCI connection request. id=%s", id)

	resp, err := getCircleCIConnection(ctx, s.client.gql, id)
	if err != nil {
		return nil, err
	}

	if resp.Vcs.GetCircleCIConnection == nil {
		return nil, fmt.Errorf("CircleCI connection not found. id=%s", id)
	}

	return resp.Vcs.GetCircleCIConnection, nil
}

// Updates the CircleCI connection with the given id.
func (s *CircleCIIntegrationService) Update(ctx context.Context, id string, name *string, apiToken *string) (*UpdateCircleCIConnectionResult, error) {
	log.Printf("update CircleCI connection request. id=%s", id)

	resp, err := updateCircleCIConnection(ctx, s.client.gql, id, name, apiToken)
	if err != nil {
		return nil, err
	}

	result := &resp.Vcs.UpdateCircleCIConnection
	log.Printf("update CircleCI connection success. id=%s", result.Id)

	return result, nil
}

// Deletes the CircleCI connection with the given id.
func (s *CircleCIIntegrationService) Delete(ctx context.Context, id string) error {
	log.Printf("delete CircleCI connection request. id=%s", id)

	resp, err := deleteCircleCIConnection(ctx, s.client.gql, id)
	if err != nil {
		return err
	}

	if !resp.Vcs.DeleteCircleCIConnection.Success {
		return fmt.Errorf("delete CircleCI connection failed. id=%s", id)
	}

	log.Printf("delete CircleCI connection success. id=%s", id)
	return nil
}
