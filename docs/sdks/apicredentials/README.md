# ApiCredentials

## Overview

Operations related to API credentials

### Available Operations

* [List](#list) - List API credentials for an Organization

## List

Retrieve all API credentials for the specified Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.apiCredentials.listApiCredentials" method="get" path="/v1/organizations/{organizationId}/api-credentials" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
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

    res, err := s.APICredentials.List(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.APICredentialsListResponseDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `organizationID`                                                       | *string*                                                               | :heavy_check_mark:                                                     | The <code>id</code> of the Organization that owns the API credentials. |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.V1APICredentialsListAPICredentialsResponse](../../models/operations/v1apicredentialslistapicredentialsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |