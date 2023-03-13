# Solarwinds Observability Client for Golang (swo-client-go)

The Solarwinds Observability Client is a Go client library for accessing the [Solarwinds Observability Api][].
The resources that are currently supported are:

* Alerts
* Dashboards
* Notification Services

## Installation ##

Currently, **swo-client-go requires Go version 1.18 or greater**.
swo-client-go is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/solarwindscloud/swo-client-go/v1
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/solarwindscloud/swo-client-go/v1"
```

and run `go get` without parameters.

Finally, to use the top-of-trunk version of this repo, use the following command:

```bash
go get github.com/solarwindscloud/swo-client-go/v1@main
```

## Usage ##
```go
// with go modules enabled (GO111MODULE=on or outside GOPATH)
import "github.com/solarwindscloud/swo-client-go/v1"
// with go modules disabled
import "github.com/solarwindscloud/swo-client-go"
```

Construct a new SWO client, then use the various services on the client to
access different parts of the SWO API. For example:

```go
client := swoClient.NewClient(apiToken, options...)

// get a single dashboard.
dashboard, err := client.DashboardService().Read(context.Background(), "[dashboard_id]")
```

The services of the client separate the API into logical groups that are directly related to the SWO API documentation at
https://api.solarwinds.com/graphql.

NOTE: The [context](https://godoc.org/context) package, can be used to pass cancelation signals and deadlines to various services of the client for handling a request. If there is no context available, then `context.Background()` can be used as a starting point.

For more sample code snippets, head over to the [example](https://github.com/solarwindscloud/swo-client-go/tree/master/example) directory.

### Authentication ###
The `swo-client-go` library handles Bearer token authentication by default using a valid auth token that must be provided to the client. The caller can provide a custom transport to the client which will allow for additional authentication methods or middleware if needed:

```go
func main() {
	// A custom transport can be used for various needs like specialized server authentication.
	transport := dev.NewUserSessionTransport(sessionId, csrfToken)
	ctx := context.Background()

	client := swoClient.NewClient(apiToken,
		swoClient.TransportOption(transport),
	)

	dashboard, err := client.DashboardService().Read(ctx, "123")
}
```

### Creating and Updating Resources ###
All structs for SWO resources use pointer values for all non-repeated fields. This allows distinguishing between unset fields and those set to a zero-value. A helper function is available `swoClient.Ptr()` to help create these pointers for common data types (e.g. string, bool, and int values). For example:

```go
// create a new private dashboard named "foo"
dashboard := &swoClient.Dashboard{
  Name:      swoClient.Ptr("my new dashboard"),
  IsPrivate: swoClient.Ptr(true),
}

resp, err := client.DashboardService().Create(ctx, dashboard)
```

### Testing code that uses `swo-client-go`
TBD

### Integration Tests ###
TBD

## Contributing ##
TBD

## Versioning ##
In general, swo-client-go follows [semver](https://semver.org/) as closely as possible for tagging releases of the package.

## License ##
This library is distributed under the Apache-style license found in the [LICENSE](./LICENSE)
file.

### Example usage
See the ./example directory for working examples of using the api.

### Issues/Bugs
Please report bugs and request enhancements in the [Issues area](https://github.com/solarwindscloud/swo-client-go/issues) of this repo.

## Requirements

- [Go](https://golang.org/doc/install) >= 1.18

## Building The Client

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Developing the Client
TBD