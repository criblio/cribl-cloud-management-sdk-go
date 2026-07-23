# Billing

## Overview

Operations related to Billing and FinOps data

### Available Operations

* [GetContractsUtilization](#getcontractsutilization) - [In development] Get contract credit utilization
* [GetCreditsTimeseries](#getcreditstimeseries) - [In development] Get credits timeseries
* [GetCreditsStats](#getcreditsstats) - [In development] Get credit balance and consumption
* [GetCreditsGrants](#getcreditsgrants) - [In development] Get credit grants

## GetContractsUtilization

**This endpoint is in development with restricted availability**. We do not recommend using this endpoint in production. Functionality might change without notice. Contact [Cribl Support](https://cribl.io/support/) to request access.

Get cumulative credit utilization data grouped by contract period for the specified Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.billing.getContractsUtilization" method="get" path="/v1/organizations/{organizationId}/billing/credits/contracts-utilization" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"github.com/criblio/cribl-cloud-management-sdk-go/types"
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

    res, err := s.Billing.GetContractsUtilization(ctx, "<id>", types.MustTimeFromString("2025-05-01T00:00:00Z"), types.MustTimeFromString("2025-06-01T00:00:00Z"), components.BillingWindowMonthly)
    if err != nil {
        log.Fatal(err)
    }
    if res.ContractsUtilizationResponseDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `organizationID`                                                     | `string`                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Organization.                             |                                                                      |
| `startingOn`                                                         | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | Inclusive start of the query date range in ISO 8601 format.          | 2025-05-01 00:00:00 +0000 UTC                                        |
| `endingBefore`                                                       | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | Exclusive end of the query date range in ISO 8601 format.            | 2025-06-01 00:00:00 +0000 UTC                                        |
| `window`                                                             | [components.BillingWindow](../../models/components/billingwindow.md) | :heavy_check_mark:                                                   | Aggregation granularity for credit utilization data.                 |                                                                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |

### Response

**[*operations.V1BillingGetContractsUtilizationResponse](../../models/operations/v1billinggetcontractsutilizationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCreditsTimeseries

**This endpoint is in development with restricted availability**. We do not recommend using this endpoint in production. Functionality might change without notice. Contact [Cribl Support](https://cribl.io/support/) to request access.

Get credit consumption timeseries data for the specified Organization, broken down by dimensions such as product and usage type.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.billing.getCreditsTimeseries" method="get" path="/v1/organizations/{organizationId}/billing/credits/timeseries" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"github.com/criblio/cribl-cloud-management-sdk-go/types"
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

    res, err := s.Billing.GetCreditsTimeseries(ctx, "<id>", types.MustTimeFromString("2025-05-01T00:00:00Z"), types.MustTimeFromString("2025-06-01T00:00:00Z"), components.BillingWindowMonthly)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditsTimeseriesResponseDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `organizationID`                                                     | `string`                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Organization.                             |                                                                      |
| `startingOn`                                                         | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | Inclusive start of the query date range in ISO 8601 format.          | 2025-05-01 00:00:00 +0000 UTC                                        |
| `endingBefore`                                                       | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | Exclusive end of the query date range in ISO 8601 format.            | 2025-06-01 00:00:00 +0000 UTC                                        |
| `window`                                                             | [components.BillingWindow](../../models/components/billingwindow.md) | :heavy_check_mark:                                                   | Aggregation granularity for timeseries data.                         |                                                                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |

### Response

**[*operations.V1BillingGetCreditsTimeseriesResponse](../../models/operations/v1billinggetcreditstimeseriesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCreditsStats

**This endpoint is in development with restricted availability**. We do not recommend using this endpoint in production. Functionality might change without notice. Contact [Cribl Support](https://cribl.io/support/) to request access.

Get credit balance totals, consumption totals, and contract date ranges for the specified Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.billing.getCreditsStats" method="get" path="/v1/organizations/{organizationId}/billing/credits/stats" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"github.com/criblio/cribl-cloud-management-sdk-go/types"
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

    res, err := s.Billing.GetCreditsStats(ctx, "<id>", types.MustTimeFromString("2025-05-01T00:00:00Z"), types.MustTimeFromString("2025-06-01T00:00:00Z"), components.BillingWindowMonthly)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditsStatsResponseDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `organizationID`                                                     | `string`                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Organization.                             |                                                                      |
| `startingOn`                                                         | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | Inclusive start of the query date range in ISO 8601 format.          | 2025-05-01 00:00:00 +0000 UTC                                        |
| `endingBefore`                                                       | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | Exclusive end of the query date range in ISO 8601 format.            | 2025-06-01 00:00:00 +0000 UTC                                        |
| `window`                                                             | [components.BillingWindow](../../models/components/billingwindow.md) | :heavy_check_mark:                                                   | Aggregation granularity for credit balance and consumption.          |                                                                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |

### Response

**[*operations.V1BillingGetCreditsStatsResponse](../../models/operations/v1billinggetcreditsstatsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCreditsGrants

**This endpoint is in development with restricted availability**. We do not recommend using this endpoint in production. Functionality might change without notice. Contact [Cribl Support](https://cribl.io/support/) to request access.

Get credit grants (purchases, rollovers, and refunds) for the specified Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.billing.getCreditsGrants" method="get" path="/v1/organizations/{organizationId}/billing/credits/grants" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"github.com/criblio/cribl-cloud-management-sdk-go/types"
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

    res, err := s.Billing.GetCreditsGrants(ctx, "<id>", types.MustTimeFromString("2025-05-01T00:00:00Z"), types.MustTimeFromString("2025-06-01T00:00:00Z"), components.BillingWindowDaily)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreditGrantsResponseDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `organizationID`                                                     | `string`                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Organization.                             |                                                                      |
| `startingOn`                                                         | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | Inclusive start of the query date range in ISO 8601 format.          | 2025-05-01 00:00:00 +0000 UTC                                        |
| `endingBefore`                                                       | [time.Time](https://pkg.go.dev/time#Time)                            | :heavy_check_mark:                                                   | Exclusive end of the query date range in ISO 8601 format.            | 2025-06-01 00:00:00 +0000 UTC                                        |
| `window`                                                             | [components.BillingWindow](../../models/components/billingwindow.md) | :heavy_check_mark:                                                   | Aggregation granularity for credit grant data.                       |                                                                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |                                                                      |

### Response

**[*operations.V1BillingGetCreditsGrantsResponse](../../models/operations/v1billinggetcreditsgrantsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |