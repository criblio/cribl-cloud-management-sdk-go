# Workspaces

## Overview

Operations related to Workspaces

### Available Operations

* [Create](#create) - Create a Workspace in the specified Organization
* [List](#list) - List all Workspaces for the specified Organization
* [Update](#update) - Update a Workspace
* [Delete](#delete) - Delete a Workspace
* [Get](#get) - Get a Workspace

## Create

Create a new Workspace in the specified Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.createWorkspace" method="post" path="/v1/organizations/{organizationId}/workspaces" -->
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

    res, err := s.Workspaces.Create(ctx, "<id>", components.WorkspaceCreateRequestDTO{
        WorkspaceID: "main",
        Alias: criblcloudmanagementsdkgo.Pointer("Production Environment"),
        Description: criblcloudmanagementsdkgo.Pointer("Main production workspace for customer data processing"),
        Tags: []string{
            "production",
            "customer-data",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `organizationID`                                                                             | *string*                                                                                     | :heavy_check_mark:                                                                           | The <code>id</code> of the Organization where you want to create the Workspace.              |
| `workspaceCreateRequestDTO`                                                                  | [components.WorkspaceCreateRequestDTO](../../models/components/workspacecreaterequestdto.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V1WorkspacesCreateWorkspaceResponse](../../models/operations/v1workspacescreateworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get a list of all Workspaces for the specified Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.listWorkspaces" method="get" path="/v1/organizations/{organizationId}/workspaces" -->
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

    res, err := s.Workspaces.List(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspacesListResponseDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `organizationID`                                                      | *string*                                                              | :heavy_check_mark:                                                    | The <code>id</code> of the Organization that contains the Workspaces. |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.V1WorkspacesListWorkspacesResponse](../../models/operations/v1workspaceslistworkspacesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.updateWorkspace" method="patch" path="/v1/organizations/{organizationId}/workspaces/{workspaceId}" -->
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

    res, err := s.Workspaces.Update(ctx, "<id>", "<id>", components.WorkspacePatchRequestDTO{
        Alias: criblcloudmanagementsdkgo.Pointer("Production Environment"),
        Description: criblcloudmanagementsdkgo.Pointer("Main production workspace for customer data processing"),
        Tags: []string{
            "production",
            "customer-data",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DefaultErrorDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `organizationID`                                                                           | *string*                                                                                   | :heavy_check_mark:                                                                         | The <code>id</code> of the Organization that contains the Workspace.                       |
| `workspaceID`                                                                              | *string*                                                                                   | :heavy_check_mark:                                                                         | The <code>id</code> of the Workspace to update.                                            |
| `workspacePatchRequestDTO`                                                                 | [components.WorkspacePatchRequestDTO](../../models/components/workspacepatchrequestdto.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V1WorkspacesUpdateWorkspaceResponse](../../models/operations/v1workspacesupdateworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Workspace in the specified Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.deleteWorkspace" method="delete" path="/v1/organizations/{organizationId}/workspaces/{workspaceId}" -->
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

    res, err := s.Workspaces.Delete(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DefaultErrorDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `organizationID`                                                     | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Organization that contains the Workspace. |
| `workspaceID`                                                        | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Workspace to delete.                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.V1WorkspacesDeleteWorkspaceResponse](../../models/operations/v1workspacesdeleteworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.getWorkspace" method="get" path="/v1/organizations/{organizationId}/workspaces/{workspaceId}" -->
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

    res, err := s.Workspaces.Get(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `organizationID`                                                     | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Organization that contains the Workspace. |
| `workspaceID`                                                        | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Workspace to get.                         |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.V1WorkspacesGetWorkspaceResponse](../../models/operations/v1workspacesgetworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |