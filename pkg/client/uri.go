package client

import (
	"context"
	"fmt"
	"log"

	"github.com/cenkalti/backoff/v5"
)

type UriService service

type CreateUriResult = createUriMutationDemDemMutationsCreateUriCreateUriResponse
type UpdateUriResult = updateUriMutationDemDemMutationsUpdateUriUpdateUriResponse
type ReadUriResult = getUriByIdEntitiesEntityQueriesByIdUri

type UriCommunicator interface {
	Create(context.Context, CreateUriInput) (*CreateUriResult, error)
	Read(context.Context, string) (*ReadUriResult, error)
	Update(context.Context, UpdateUriInput) error
	Delete(context.Context, string) error
}

func newUriService(c *Client) *UriService {
	return &UriService{c}
}

// Creates a new Uri entity with the given input.
func (as *UriService) Create(ctx context.Context, input CreateUriInput) (*CreateUriResult, error) {
	log.Printf("create Uri request. name=%s, url=%s", input.Name, input.IpOrDomain)

	resp, err := createUriMutation(ctx, as.client.gql, input)
	if err != nil {
		return nil, err
	}

	result := resp.Dem.CreateUri
	log.Printf("create Uri success. id=%s", result.Id)

	return &result, nil
}

// Returns the Uri entity with the given Id.
func (as *UriService) Read(ctx context.Context, id string) (*ReadUriResult, error) {
	log.Printf("read uri request. id=%s", id)

	operation := func() (getUriByIdEntitiesEntityQueriesByIdEntity, error) {
		resp, err := getUriById(ctx, as.client.gql, id)

		if err != nil {
			return nil, err
		}

		if resp.Entities.ById == nil {
			return nil, backoff.RetryAfter(1)
		}

		return *resp.Entities.ById, nil
	}

	uriPtr, err := backoff.Retry(ctx, operation, backoff.WithBackOff(backoff.NewExponentialBackOff()), backoff.WithMaxElapsedTime(5))

	if err != nil {
		return nil, err
	}

	if uri, ok := uriPtr.(*ReadUriResult); !ok {
		return nil, fmt.Errorf("unexpected type %T", uri)
	} else {
		return uri, nil
	}
}

// Updates the Uri with input for the given id.
func (as *UriService) Update(ctx context.Context, input UpdateUriInput) error {
	log.Printf("update uri request. id=%s", input.Id)

	if _, err := updateUriMutation(ctx, as.client.gql, input); err != nil {
		return err
	}

	log.Printf("update uri success. id=%s", input.Id)
	return nil
}

// Deletes the Uri with the given id.
func (as *UriService) Delete(ctx context.Context, id string) error {
	log.Printf("delete uri request. id=%s", id)

	if _, err := deleteUriMutation(ctx, as.client.gql, DeleteUriInput{id}); err != nil {
		return err
	}

	log.Printf("delete uri success. id=%s", id)
	return nil
}
