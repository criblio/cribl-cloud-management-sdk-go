# CriblMgmtPlane SDK

## Overview

Cribl Cloud Management API: Lorem Ipsum

### Available Operations

* [DummyServiceStatus](#dummyservicestatus) - Service status

## DummyServiceStatus

Service status

### Example Usage

```go
package main

import(
	"context"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcloudmanagementsdkgo.New(
        "https://api.example.com",
        criblcloudmanagementsdkgo.WithSecurity(os.Getenv("CRIBLMGMTPLANE_BEARER_AUTH")),
    )

    res, err := s.DummyServiceStatus(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.DummyServiceStatusResponse](../../models/operations/dummyservicestatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |