package client

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Khan/genqlient/graphql"
	log "github.com/sirupsen/logrus"
)

const (
	defaultBaseURL        = "https://api.dc-01.cloud.solarwinds.com/graphql"
	defaultMediaType      = "application/json"
	defaultRequestTimeout = 30 * time.Second
	clientIdentifier      = "Swo-Api-Go"
	requestIdentifier     = "X-Request-Id"
)

// ServiceAccessor defines an interface for talking to via domain-specific service constructs
type ServiceAccessor interface {
	AlertsService() AlertsCommunicator
	DashboardsService() DashboardsCommunicator
	NotificationsService() NotificationsCommunicator
	UriService() UriCommunicator
	WebsiteService() WebsiteCommunicator
}

// Client implements ServiceAccessor
type Client struct {
	// Option settings
	baseURL        *url.URL
	debugMode      bool
	requestTimeout time.Duration
	userAgent      string
	transport      http.RoundTripper

	// GraphQL client
	gql graphql.Client

	// Service accessors
	alertsService        AlertsCommunicator
	dashboardsService    DashboardsCommunicator
	notificationsService NotificationsCommunicator
	uriService           UriCommunicator
	websiteService       WebsiteCommunicator
}

// Each service derives from the service type.
type service struct {
	client *Client
}

// Returns a new SWO API client with functional override options.
// * BaseUrlOption
// * DebugOption
// * TransportOption
// * UserAgentOption
// * RequestTimeoutOption
func NewClient(apiToken string, opts ...ClientOption) *Client {
	baseURL, err := url.Parse(defaultBaseURL)

	if err != nil {
		log.Error(err)
		return nil
	}

	swoClient := &Client{
		baseURL:        baseURL,
		requestTimeout: defaultRequestTimeout,
	}

	// Set any user options that were provided.
	for _, opt := range opts {
		err = opt(swoClient)
		if err != nil {
			log.Error(fmt.Errorf("client option error. fallback to default value: %s", err))
		}
	}

	// Use the api token transport if one wasn't provided.
	if swoClient.transport == nil {
		swoClient.transport = &apiTokenAuthTransport{
			apiToken: apiToken,
			client:   swoClient,
		}
	}

	if swoClient.debugMode {
		log.SetLevel(log.TraceLevel)
		log.Info("swoclient: debugMode set to true.")
	}

	swoClient.gql = graphql.NewClient(swoClient.baseURL.String(), &http.Client{
		Timeout:   swoClient.requestTimeout,
		Transport: swoClient.transport,
	})

	return swoClient
}

// A subset of the API that deals with Alerts.
func (c *Client) AlertsService() AlertsCommunicator {
	if c.alertsService == nil {
		c.alertsService = NewAlertsService(c)
	}

	return c.alertsService
}

// A subset of the API that deals with Dashboards.
func (c *Client) DashboardsService() DashboardsCommunicator {
	if c.dashboardsService == nil {
		c.dashboardsService = NewDashboardsService(c)
	}

	return c.dashboardsService
}

// A subset of the API that deals with Notifications.
func (c *Client) NotificationsService() NotificationsCommunicator {
	if c.notificationsService == nil {
		c.notificationsService = NewNotificationsService(c)
	}

	return c.notificationsService
}

// A subset of the API that deals with Uris.
func (c *Client) UriService() UriCommunicator {
	if c.uriService == nil {
		c.uriService = NewUriService(c)
	}

	return c.uriService
}

// A subset of the API that deals with Websites.
func (c *Client) WebsiteService() WebsiteCommunicator {
	if c.websiteService == nil {
		c.websiteService = NewWebsiteService(c)
	}

	return c.websiteService
}

// Returns the string that will be placed in the User-Agent header. It ensures
// that any caller-set string has the client name and version appended to it.
func (c *Client) completeUserAgentString() string {
	if c.userAgent == "" {
		return clientIdentifier
	}
	return fmt.Sprintf("%s:%s", c.userAgent, clientIdentifier)
}
