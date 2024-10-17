package client

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type ApiTokenService service

type CreateApiTokenResult = createTokenMutationCreateTokenCreateTokenResponseToken
type ReadApiTokenResult = getApiTokenByIdUserAuthenticatedUserCurrentOrganizationTokensToken

type ApiTokenCommunicator interface {
	Create(context.Context, CreateTokenInput) (*CreateApiTokenResult, error)
	Read(context.Context, string) (*ReadApiTokenResult, error)
	Update(context.Context, UpdateTokenInput) error
	Delete(context.Context, string) error
}

func newApiTokenService(c *Client) *ApiTokenService {
	return &ApiTokenService{c}
}

// Creates a new ApiToken entity with the given input.
func (as *ApiTokenService) Create(ctx context.Context, input CreateTokenInput) (*CreateApiTokenResult, error) {
	log.Printf("create apiToken request. name=%s", input.Name)

	resp, err := doMutate(
		func() (*createTokenMutationResponse, error) {
			return createTokenMutation(ctx, as.client.gql, input)
		},
		func(resp *createTokenMutationResponse) error {
			if !resp.CreateToken.Success {
				return mutateError("update apiToken failed",
					resp.CreateToken.Code,
					resp.CreateToken.Message)
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	result := resp.CreateToken.Token
	log.Printf("create apiToken success. id=%s", result.Id)

	return result, nil
}

// Returns the ApiToken entity with the given Id.
func (as *ApiTokenService) Read(ctx context.Context, id string) (*ReadApiTokenResult, error) {
	log.Printf("read apiToken request. id=%s", id)

	resp, err := getApiTokenById(ctx, as.client.gql, id)
	if err != nil {
		return nil, err
	}

	if len(resp.User.CurrentOrganization.Tokens) == 0 {
		return nil, errors.New(fmt.Sprintf("api token not found. id=%s", id))
	}

	return &resp.User.CurrentOrganization.Tokens[0], nil
}

// Updates the ApiToken with input for the given id.
func (as *ApiTokenService) Update(ctx context.Context, input UpdateTokenInput) error {
	log.Printf("update apiToken request. id=%s", input.Id)

	if _, err := doMutate(
		func() (*updateTokenMutationResponse, error) {
			return updateTokenMutation(ctx, as.client.gql, input)
		},
		func(resp *updateTokenMutationResponse) error {
			if !resp.UpdateToken.Success {
				return mutateError("update apiToken failed",
					resp.UpdateToken.Code,
					resp.UpdateToken.Message)
			}
			return nil
		}); err != nil {
		return err
	}

	log.Printf("update apiToken success. id=%s", input.Id)
	return nil
}

// Deletes the ApiToken with the given id.
func (as *ApiTokenService) Delete(ctx context.Context, id string) error {
	log.Printf("delete apiToken request. id=%s", id)

	if _, err := doMutate(
		func() (*deleteTokenMutationResponse, error) {
			return deleteTokenMutation(ctx, as.client.gql, DeleteTokenInput{Id: id})
		},
		func(resp *deleteTokenMutationResponse) error {
			if !resp.DeleteToken.Success {
				return mutateError("update apiToken failed",
					resp.DeleteToken.Code,
					resp.DeleteToken.Message)
			}
			return nil
		}); err != nil {
		return err
	}

	log.Printf("delete apiToken success. id=%s", id)
	return nil
}
