package client

import (
	"context"
	"fmt"
	"log"
)

type WebsiteService service

type CreateWebsiteResult = createWebsiteMutationDemDemMutationsCreateWebsiteCreateWebsiteSuccess
type UpdateWebsiteResult = updateWebsiteMutationResponse
type ReadWebsiteResult = getWebsiteByIdEntitiesEntityQueriesByIdWebsite

type WebsiteCommunicator interface {
	Create(context.Context, CreateWebsiteInput) (*CreateWebsiteResult, error)
	Read(context.Context, string) (*ReadWebsiteResult, error)
	Update(context.Context, UpdateWebsiteInput) error
	Delete(context.Context, string) error
}

func newWebsiteService(c *Client) *WebsiteService {
	return &WebsiteService{c}
}

// Creates a new website entity with the given input.
func (as *WebsiteService) Create(ctx context.Context, input CreateWebsiteInput) (*CreateWebsiteResult, error) {
	log.Printf("create website request. name=%s, url=%s", input.Name, input.Url)

	resp, err := createWebsiteMutation(ctx, as.client.gql, input)

	if err != nil {
		return nil, err
	}

	website := resp.Dem.CreateWebsite
	log.Printf("create website success. id=%s", website.Id)

	return &website, nil
}

// Returns the website entity with the given Id.
func (as *WebsiteService) Read(ctx context.Context, id string) (*ReadWebsiteResult, error) {
	log.Printf("read website request. Id: %s", id)

	resp, err := getWebsiteById(ctx, as.client.gql, id)
	if err != nil {
		return nil, err
	}

	if resp.Entities.ById == nil {
		return nil, ErrEntityIdNil
	}

	websitePtr := *resp.Entities.ById

	website, ok := websitePtr.(*ReadWebsiteResult)
	if !ok {
		return nil, fmt.Errorf("unexpected type %T", website)
	}

	return website, nil
}

// Updates the website with input for the given id.
func (as *WebsiteService) Update(ctx context.Context, input UpdateWebsiteInput) error {
	log.Printf("update website request. id=%s", input.Id)

	if _, err := updateWebsiteMutation(ctx, as.client.gql, input); err != nil {
		return err
	}

	log.Printf("update website success. id=%s", input.Id)
	return nil
}

// Deletes the website with the given id.
func (as *WebsiteService) Delete(ctx context.Context, id string) error {
	log.Printf("delete website request. id=%s", id)

	if _, err := deleteWebsiteMutation(ctx, as.client.gql, DeleteWebsiteInput{id}); err != nil {
		return err
	}

	log.Printf("delete website success. id=%s", id)
	return nil
}
