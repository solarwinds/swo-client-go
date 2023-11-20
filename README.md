# Solarwinds Observability Client for Golang
## swo-client-go ##

The Solarwinds Observability Client is a Go client library for accessing the [Solarwinds Observability Api]().
The resources that are currently supported are:

* Alerts
* Dashboards
* Notification Services
* Websites
* Uris

## Installation ##
Currently, **swo-client-go requires Go version 1.18 or greater**.
swo-client-go is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/solarwinds/swo-client-go/v1
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/solarwinds/swo-client-go/v1"
```

and run `go get` without parameters.

Finally, to use the top-of-trunk version of this repo, use the following command:

```bash
go get github.com/solarwinds/swo-client-go/v1@main
```

## Usage ##
```go
// with go modules enabled (GO111MODULE=on or outside GOPATH)
import "github.com/solarwinds/swo-client-go/v1"
// with go modules disabled
import "github.com/solarwinds/swo-client-go"
```

Construct a new SWO client, then use the various services on the client to
access different parts of the SWO API. For example:

```go
client := swoClient.New(apiToken, options...)

// get a single dashboard.
dashboard, err := client.DashboardService().Read(context.Background(), "[dashboard_id]")
```

The services of the client separate the API into logical groups that are directly related to the SWO API documentation at
https://api.solarwinds.com/graphql.

NOTE: The [context](https://godoc.org/context) package, can be used to pass cancelation signals and deadlines to various services of the client for handling a request. If there is no context available, then `context.Background()` can be used as a starting point.

For more sample code snippets, head over to the [example](https://github.com/solarwinds/swo-client-go/tree/master/example) directory.

### Authentication ###
The `swo-client-go` library handles Bearer token authentication by default using a valid api token that must be provided to the client. The caller can provide a custom transport to the client which will allow for additional authentication methods or middleware if needed:

```go
func main() {
  // A custom transport can be used for various needs like specialized server authentication.
  transport := NewCustomTransport(myOptions)
  ctx := context.Background()

  client, err := swoClient.New(apiToken,
    swoClient.TransportOption(transport),
  )

  dashboard, err := client.DashboardService().Read(ctx, "123")
}
```

### Creating and Updating Resources ###
All structs for SWO resources use pointer values for all non-repeated fields. This allows distinguishing between unset fields and those set to a zero-value. A helper function is available `Ptr()` to help create these pointers for common data types (e.g. string, bool, and int values). For example:

```go
// create a new private dashboard named "My New Dashboard"
input := &CreateDashboardInput{
  Name:      swo.Ptr("My New Dashboard"),
  IsPrivate: swo.Ptr(true),
}

dashboard, err := client.DashboardService().Create(ctx, input)
```

## Versioning ##
In general, swo-client-go follows [semver](https://semver.org/) as closely as possible for tagging releases of the package.

## License ##
This library is distributed under the Apache-style license found in the [LICENSE](./LICENSE)
file.

### Example usage
See the ./example directory for working examples of using the api.

### Issues/Bugs
Please report bugs and request enhancements in the [Issues area](https://github.com/solarwinds/swo-client-go/issues) of this repo.

## Requirements
- [Go](https://golang.org/doc/install) >= 1.18

## Building The Client
1. Clone the repository
1. Enter the repository directory
1. Build the client using the Go `install` command:

```shell
go install
```

## Developing the Client
The SWO api is a GraphQL endpoint that allows for simple extensions to the the client library. There are a few steps required to generate code from new schemas, queries, and mutations. The tool that performs the code generation is called [genqlient](github.com/Khan/genqlient) and is part of the tools in this project.

To add new queries and mutations simply add your GQL to the ./graphql/ folder in this project. There are a few existing services there as an example. The GQL is separated by service in the same way the client api is. If adding support for a new service to the project you must also include the new files in the `./graphql/genqlient.yaml` configuration so the genqlient tool knows to generate code for them.

Once you are ready to generate new code simply run the tool against the `./graphql/` folder.
```shell
cd graphql
go run github.com/Khan/genqlient
```
The resulting Go code is output to the `./pkg/client/genqlient_generated.go` file and can now be used for development.
