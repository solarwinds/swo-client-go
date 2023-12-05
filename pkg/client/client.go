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
	LogFilterService() LogFilterCommunicator
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
	apiTokenService      ApiTokenCommunicator
	dashboardsService    DashboardsCommunicator
	logFilterService     LogFilterCommunicator
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
	c.alertsService = newAlertsService(c)
	c.apiTokenService = newApiTokenService(c)
	c.dashboardsService = newDashboardsService(c)
	c.logFilterService = newLogFilterService(c)
	c.notificationsService = newNotificationsService(c)
	c.uriService = newUriService(c)
	c.websiteService = newWebsiteService(c)

	// We will keep this available in case more complex initialization is needed in the future.
	return nil
}

// A subset of the API that deals with Alerts.
func (c *Client) AlertsService() AlertsCommunicator {
	return c.alertsService
}

// A subset of the API that deals with ApiTokens.
func (c *Client) ApiTokenService() ApiTokenCommunicator {
	return c.apiTokenService
}

// A subset of the API that deals with Dashboards.
func (c *Client) DashboardsService() DashboardsCommunicator {
	return c.dashboardsService
}

// A subset of the API that deals with LogFilters.
func (c *Client) LogFilterService() LogFilterCommunicator {
	return c.logFilterService
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
