package client

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type LogFilterService service

type CreateLogFilterResult = createLogFilterCreateExclusionFilterCreateExclusionFilterResponseExclusionFilter
type ReadLogFilterResult = getLogFilterByIdGetExclusionFilter

type LogFilterCommunicator interface {
	Create(context.Context, CreateExclusionFilterInput) (*CreateLogFilterResult, error)
	Read(context.Context, string) (*ReadLogFilterResult, error)
	Update(context.Context, UpdateExclusionFilterInput) error
	Delete(context.Context, string) error
}

func newLogFilterService(c *Client) *LogFilterService {
	return &LogFilterService{c}
}

// Creates a new LogFilter entity with the given input.
func (as *LogFilterService) Create(ctx context.Context, input CreateExclusionFilterInput) (*CreateLogFilterResult, error) {
	log.Printf("create LogFilter request. name=%s", input.Name)

	resp, err := createLogFilter(ctx, as.client.gql, input)
	if err != nil {
		return nil, err
	}

	createFilter := resp.CreateExclusionFilter
	if createFilter.Success == false {
		err := fmt.Sprintf("create LogFilter failed. code=%s message=%s", createFilter.Code, createFilter.Message)
		log.Print(err)
		return nil, errors.New(err)
	}

	filter := createFilter.ExclusionFilter
	log.Printf("create LogFilter success. id=%s", filter.Id)
	return filter, nil
}

// Returns the LogFilter entity with the given Id.
func (as *LogFilterService) Read(ctx context.Context, id string) (*ReadLogFilterResult, error) {
	log.Printf("read logFilter request. id=%s", id)

	resp, err := getLogFilterById(ctx, as.client.gql, GetExclusionFilterInput{Id: id})
	if err != nil {
		return nil, err
	}

	return resp.GetExclusionFilter, nil
}

// Updates the LogFilter with input for the given id.
func (as *LogFilterService) Update(ctx context.Context, input UpdateExclusionFilterInput) error {
	log.Printf("update logFilter request. id=%s", input.Id)

	if _, err := updateLogFilter(ctx, as.client.gql, input); err != nil {
		return err
	}

	log.Printf("update logFilter success. id=%s", input.Id)
	return nil
}

// Deletes the LogFilter with the given id.
func (as *LogFilterService) Delete(ctx context.Context, id string) error {
	log.Printf("delete logFilter request. id=%s", id)

	if _, err := deleteLogFilter(ctx, as.client.gql, DeleteExclusionFilterInput{id}); err != nil {
		return err
	}

	log.Printf("delete logFilter success. id=%s", id)
	return nil
}
