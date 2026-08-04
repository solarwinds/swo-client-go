package client

import (
	"context"
	"fmt"
	"log"
	"strings"
)

// Error implements the error interface for the genqlient-generated DashboardValidationError
// type (see the DashboardValidationError fragment in graphql/dashboards.graphql), so it can
// be used directly wherever a single field-level validation failure needs to be reported.
func (e DashboardValidationError) Error() string {
	msg := fmt.Sprintf("field=%s code=%s message=%s", e.Field, e.Code, e.Message)
	if e.Constraints == nil || (e.Constraints.Min == nil && e.Constraints.Max == nil) {
		return msg
	}

	constraint := ""
	if e.Constraints.Min != nil {
		constraint += fmt.Sprintf("min=%v", *e.Constraints.Min)
	}
	if e.Constraints.Max != nil {
		if constraint != "" {
			constraint += ", "
		}
		constraint += fmt.Sprintf("max=%v", *e.Constraints.Max)
	}
	return fmt.Sprintf("%s (%s)", msg, constraint)
}

// DashboardMutationError is returned when a create/update/delete dashboard mutation
// fails. It implements the error interface, so existing callers keep working, but also
// exposes the structured code, message and per-field validation errors so callers that
// need them can access them programmatically (e.g. via errors.As) instead of parsing
// the error string.
type DashboardMutationError struct {
	// Op describes the failed operation, e.g. "create dashboard".
	Op string
	// Code is the machine-readable, server-provided error code.
	Code string
	// Message is the human-readable, server-provided error message.
	Message string
	// ValidationErrors holds any field-level validation failures, if the server reported them.
	ValidationErrors []DashboardValidationError
}

func (e *DashboardMutationError) Error() string {
	msg := fmt.Sprintf("%s failed. code: %s message: %s", e.Op, e.Code, e.Message)
	if len(e.ValidationErrors) == 0 {
		return msg
	}

	details := make([]string, len(e.ValidationErrors))
	for i, ve := range e.ValidationErrors {
		details[i] = ve.Error()
	}
	return fmt.Sprintf("%s validationErrors: [%s]", msg, strings.Join(details, "; "))
}

type DashboardsService service

type CreateDashboardResult = createDashboardCreateDashboardCreateDashboardResponseDashboard
type CreateDashboardLayout = createDashboardCreateDashboardCreateDashboardResponseDashboardLayout
type CreateDashboardWidget = createDashboardCreateDashboardCreateDashboardResponseDashboardWidgetsWidget

type ReadDashboardResult = getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboard
type ReadDashboardLayout = getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardLayout
type ReadDashboardWidget = getDashboardByIdDashboardsDashboardQueriesByIdOrSystemReferenceDashboardWidgetsWidget

type UpdateDashboardResult = updateDashboardUpdateDashboardUpdateDashboardResponseDashboard
type UpdateDashboardLayout = updateDashboardUpdateDashboardUpdateDashboardResponseDashboardLayout
type UpdateDashboardWidget = updateDashboardUpdateDashboardUpdateDashboardResponseDashboardWidgetsWidget

type DashboardsCommunicator interface {
	Create(context.Context, CreateDashboardInput) (*CreateDashboardResult, error)
	Read(context.Context, string) (*ReadDashboardResult, error)
	Update(context.Context, UpdateDashboardInput) (*UpdateDashboardResult, error)
	Delete(context.Context, string) error
}

func newDashboardsService(c *Client) *DashboardsService {
	return &DashboardsService{c}
}

// newCreateDashboardValidationErrors unwraps the genqlient createDashboard response items
// down to the shared, genqlient-generated DashboardValidationError type.
func newCreateDashboardValidationErrors(
	errs []createDashboardCreateDashboardCreateDashboardResponseValidationErrorsDashboardValidationError,
) []DashboardValidationError {
	if len(errs) == 0 {
		return nil
	}

	out := make([]DashboardValidationError, len(errs))
	for i, ve := range errs {
		out[i] = ve.DashboardValidationError
	}
	return out
}

// newUpdateDashboardValidationErrors unwraps the genqlient updateDashboard response items
// down to the shared, genqlient-generated DashboardValidationError type.
func newUpdateDashboardValidationErrors(
	errs []updateDashboardUpdateDashboardUpdateDashboardResponseValidationErrorsDashboardValidationError,
) []DashboardValidationError {
	if len(errs) == 0 {
		return nil
	}

	out := make([]DashboardValidationError, len(errs))
	for i, ve := range errs {
		out[i] = ve.DashboardValidationError
	}
	return out
}

// Creates a new dashboard.
func (service *DashboardsService) Create(ctx context.Context, input CreateDashboardInput) (*CreateDashboardResult, error) {
	log.Printf("create dashboard request. name: %s", input.Name)

	resp, err := doMutate(
		func() (*createDashboardResponse, error) {
			return createDashboard(ctx, service.client.gql, input)
		},
		func(resp *createDashboardResponse) error {
			if !resp.CreateDashboard.Success {
				return &DashboardMutationError{
					Op:               "create dashboard",
					Code:             resp.CreateDashboard.Code,
					Message:          resp.CreateDashboard.Message,
					ValidationErrors: newCreateDashboardValidationErrors(resp.CreateDashboard.ValidationErrors),
				}
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	dashboard := resp.CreateDashboard.Dashboard
	log.Printf("create dashboard success. id: %s", dashboard.Id)

	return dashboard, nil
}

// Returns the dashboard identified by the given Id.
func (service *DashboardsService) Read(ctx context.Context, id string) (*ReadDashboardResult, error) {
	log.Printf("read dashboard request. id: %s", id)

	resp, err := getDashboardById(ctx, service.client.gql, id)

	if err != nil {
		return nil, err
	}

	dashboard := resp.Dashboards.ByIdOrSystemReference

	log.Printf("read dashboard success. name: %s", dashboard.Id)

	return dashboard, nil
}

// Updates the dashboard.
func (service *DashboardsService) Update(ctx context.Context, input UpdateDashboardInput) (*UpdateDashboardResult, error) {
	log.Printf("update dashboard request. id: %s", input.Id)

	resp, err := doMutate(
		func() (*updateDashboardResponse, error) {
			return updateDashboard(ctx, service.client.gql, input)
		},
		func(resp *updateDashboardResponse) error {
			if !resp.UpdateDashboard.Success {
				return &DashboardMutationError{
					Op:               "update dashboard",
					Code:             resp.UpdateDashboard.Code,
					Message:          resp.UpdateDashboard.Message,
					ValidationErrors: newUpdateDashboardValidationErrors(resp.UpdateDashboard.ValidationErrors),
				}
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	log.Printf("update dashboard success. id: %s", input.Id)

	return resp.UpdateDashboard.Dashboard, nil
}

// Deletes the dashboard with the given id.
func (service *DashboardsService) Delete(ctx context.Context, id string) error {
	log.Printf("delete dashboard request. id: %s", id)

	_, err := doMutate(
		func() (*deleteDashboardResponse, error) {
			return deleteDashboard(ctx, service.client.gql, DeleteDashboardInput{
				Id: id,
			})
		},
		func(resp *deleteDashboardResponse) error {
			if !resp.DeleteDashboard.Success {
				return mutateError("delete dashboard failed",
					resp.DeleteDashboard.Code,
					resp.DeleteDashboard.Message)
			}
			return nil
		})

	if err != nil {
		return err
	}

	log.Printf("delete dashboard success. id: %s", id)

	return nil
}
