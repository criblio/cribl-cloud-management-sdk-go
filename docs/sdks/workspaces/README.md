# Workspaces
(*Workspaces*)

## Overview

Operations related to workspaces

### Available Operations

* [Create](#create) - Create a new workspace
* [List](#list) - List all workspaces for an organization
* [Update](#update) - Update an existing workspace
* [Delete](#delete) - Delete a workspace
* [Get](#get) - Get a specific workspace by ID

## Create

Create a new workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.createWorkspace" method="post" path="/v1/organizations/{organizationId}/workspaces" -->
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
| `organizationID`                                                                             | *string*                                                                                     | :heavy_check_mark:                                                                           | Organization identifier                                                                      |
| `workspaceCreateRequestDTO`                                                                  | [components.WorkspaceCreateRequestDTO](../../models/components/workspacecreaterequestdto.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.V1WorkspacesCreateWorkspaceResponse](../../models/operations/v1workspacescreateworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List all workspaces for an organization

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.listWorkspaces" method="get" path="/v1/organizations/{organizationId}/workspaces" -->
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `organizationID`                                         | *string*                                                 | :heavy_check_mark:                                       | Organization identifier                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.V1WorkspacesListWorkspacesResponse](../../models/operations/v1workspaceslistworkspacesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an existing workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.updateWorkspace" method="patch" path="/v1/organizations/{organizationId}/workspaces/{workspaceId}" -->
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
    if res.WorkspaceSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `organizationID`                                                                           | *string*                                                                                   | :heavy_check_mark:                                                                         | Organization identifier                                                                    |
| `workspaceID`                                                                              | *string*                                                                                   | :heavy_check_mark:                                                                         | Workspace identifier                                                                       |
| `workspacePatchRequestDTO`                                                                 | [components.WorkspacePatchRequestDTO](../../models/components/workspacepatchrequestdto.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.V1WorkspacesUpdateWorkspaceResponse](../../models/operations/v1workspacesupdateworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.deleteWorkspace" method="delete" path="/v1/organizations/{organizationId}/workspaces/{workspaceId}" -->
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

    res, err := s.Workspaces.Delete(ctx, "<id>", "<id>")
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
| `organizationID`                                         | *string*                                                 | :heavy_check_mark:                                       | Organization identifier                                  |
| `workspaceID`                                            | *string*                                                 | :heavy_check_mark:                                       | Workspace identifier                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.V1WorkspacesDeleteWorkspaceResponse](../../models/operations/v1workspacesdeleteworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get a specific workspace by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.workspaces.getWorkspace" method="get" path="/v1/organizations/{organizationId}/workspaces/{workspaceId}" -->
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `organizationID`                                         | *string*                                                 | :heavy_check_mark:                                       | Organization identifier                                  |
| `workspaceID`                                            | *string*                                                 | :heavy_check_mark:                                       | Workspace identifier                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.V1WorkspacesGetWorkspaceResponse](../../models/operations/v1workspacesgetworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |