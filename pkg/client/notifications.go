package client

import (
	"context"
	"log"
)

type NotificationsService service

type CreateNotificationInput = CreateNotificationServiceConfigurationInput
type CreateNotificationResult = createNotificationCreateNotificationServiceConfigurationCreateNotificationServiceConfigurationResponseConfigurationNotificationService

type ReadNotificationResult = getNotificationUserAuthenticatedUserCurrentOrganizationNotificationServiceConfigurationNotificationService

type UpdateNotificationInput = UpdateNotificationServiceConfigurationInput
type UpdateNotificationResult = updateNotificationUpdateNotificationServiceConfigurationUpdateNotificationServiceConfigurationResponseConfigurationNotificationService

type NotificationsCommunicator interface {
	Create(context.Context, CreateNotificationInput) (*CreateNotificationResult, error)
	Read(context.Context, string, string) (*ReadNotificationResult, error)
	Update(context.Context, UpdateNotificationInput) (*UpdateNotificationResult, error)
	Delete(context.Context, string) error
}

func newNotificationsService(c *Client) *NotificationsService {
	return &NotificationsService{c}
}

// Creates a new notification.
func (service *NotificationsService) Create(ctx context.Context, input CreateNotificationInput) (*CreateNotificationResult, error) {
	log.Printf("create notification request. title: %s", input.Title)

	resp, err := doMutate(
		func() (*createNotificationResponse, error) {
			return createNotification(ctx, service.client.gql, input)
		},
		func(resp *createNotificationResponse) error {
			if !resp.CreateNotificationServiceConfiguration.Success {
				return mutateError("create notification failed",
					resp.CreateNotificationServiceConfiguration.Code,
					resp.CreateNotificationServiceConfiguration.Message)
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	notification := resp.CreateNotificationServiceConfiguration.Configuration
	log.Printf("create notifications success. id: %s", notification.Id)

	return notification, nil
}

// Returns the notification identified by the given Id.
func (service *NotificationsService) Read(ctx context.Context, id string, notificationType string) (*ReadNotificationResult, error) {
	log.Printf("read notification request. id: %s", id)

	resp, err := getNotification(ctx, service.client.gql, id, notificationType)

	if err != nil {
		return nil, err
	}

	notification := resp.User.CurrentOrganization.NotificationServiceConfiguration

	log.Printf("read notification success. title: %s", notification.Title)

	return &notification, nil
}

// Updates the notification.
func (service *NotificationsService) Update(ctx context.Context, input UpdateNotificationInput) (*UpdateNotificationResult, error) {
	log.Printf("update notification request. id: %s", input.Id)

	resp, err := doMutate(
		func() (*updateNotificationResponse, error) {
			return updateNotification(ctx, service.client.gql, input)
		},
		func(resp *updateNotificationResponse) error {
			if !resp.UpdateNotificationServiceConfiguration.Success {
				return mutateError("update notification failed",
					resp.UpdateNotificationServiceConfiguration.Code,
					resp.UpdateNotificationServiceConfiguration.Message)
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	log.Printf("update notification success. id: %s", input.Id)

	return resp.UpdateNotificationServiceConfiguration.Configuration, nil
}

// Deletes the notification with the given id.
func (service *NotificationsService) Delete(ctx context.Context, id string) error {
	log.Printf("delete notification request. id: %s", id)

	_, err := doMutate(
		func() (*deleteNotificationResponse, error) {
			return deleteNotification(ctx, service.client.gql, DeleteNotificationServiceConfigurationInput{
				Id: id,
			})
		},
		func(resp *deleteNotificationResponse) error {
			if !resp.DeleteNotificationServiceConfiguration.Success {
				return mutateError("delete notification failed",
					resp.DeleteNotificationServiceConfiguration.Code,
					resp.DeleteNotificationServiceConfiguration.Message)
			}
			return nil
		})

	if err != nil {
		return err
	}

	log.Printf("delete notification success. id: %s", id)

	return nil
}
