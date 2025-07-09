package client

import (
	"context"
	"log"
)

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

// Creates a new dashboard.
func (service *DashboardsService) Create(ctx context.Context, input CreateDashboardInput) (*CreateDashboardResult, error) {
	log.Printf("create dashboard request. name: %s", input.Name)

	resp, err := doMutate(
		func() (*createDashboardResponse, error) {
			return createDashboard(ctx, service.client.gql, input)
		},
		func(resp *createDashboardResponse) error {
			if !resp.CreateDashboard.Success {
				return mutateError("create dashboard failed",
					resp.CreateDashboard.Code,
					resp.CreateDashboard.Message)
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
				return mutateError("update dashboard failed",
					resp.UpdateDashboard.Code,
					resp.UpdateDashboard.Message)
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
