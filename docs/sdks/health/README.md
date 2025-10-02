# Health
(*Health*)

## Overview

Operations related to application health status

### Available Operations

* [Get](#get) - Get the health status of the application

## Get

Get the health status of the application

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealthStatus" method="get" path="/" -->
```go
package main

import(
	"context"
	"os"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcloudmanagementsdkgo.New(
        criblcloudmanagementsdkgo.WithSecurity(components.Security{
            ClientOauth: &components.SchemeClientOauth{
                ClientID: os.Getenv("CRIBLMGMTPLANE_CLIENT_ID"),
                ClientSecret: os.Getenv("CRIBLMGMTPLANE_CLIENT_SECRET"),
                TokenURL: os.Getenv("CRIBLMGMTPLANE_TOKEN_URL"),
                Audience: "https://api.cribl.cloud",
            },
        }),
    )

    res, err := s.Health.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetHealthStatusResponse](../../models/operations/gethealthstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |