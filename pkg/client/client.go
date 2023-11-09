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

var (
	serviceInitError = func(serviceName string) error {
		return fmt.Errorf("could not instantiate service: name=%s", serviceName)
	}
)

// Returns a new SWO API client with functional override options.
// * BaseUrlOption
// * DebugOption
// * TransportOption
// * UserAgentOption
// * RequestTimeoutOption
func New(apiToken string, opts ...ClientOption) (*Client, error) {
	baseURL, err := url.Parse(defaultBaseURL)

	if err != nil {
		return nil, err
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

	if err = initServices(swoClient); err != nil {
		return nil, err
	}

	return swoClient, nil
}

func initServices(c *Client) error {
	if c.alertsService = newAlertsService(c); c.alertsService == nil {
		return serviceInitError("AlertsService")
	}
	if c.dashboardsService = newDashboardsService(c); c.dashboardsService == nil {
		return serviceInitError("DashboardsService")
	}
	if c.notificationsService = newNotificationsService(c); c.notificationsService == nil {
		return serviceInitError("NotificationsService")
	}
	if c.uriService = newUriService(c); c.uriService == nil {
		return serviceInitError("UriService")
	}
	if c.websiteService = newWebsiteService(c); c.websiteService == nil {
		return serviceInitError("WebsiteService")
	}

	return nil
}

// A subset of the API that deals with Alerts.
func (c *Client) AlertsService() AlertsCommunicator {
	return c.alertsService
}

// A subset of the API that deals with Dashboards.
func (c *Client) DashboardsService() DashboardsCommunicator {
	return c.dashboardsService
}

// A subset of the API that deals with Notifications.
func (c *Client) NotificationsService() NotificationsCommunicator {
	return c.notificationsService
}

// A subset of the API that deals with Uris.
func (c *Client) UriService() UriCommunicator {
	return c.uriService
}

// A subset of the API that deals with Websites.
func (c *Client) WebsiteService() WebsiteCommunicator {
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
