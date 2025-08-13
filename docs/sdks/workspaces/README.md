# Workspaces
(*Workspaces*)

## Overview

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
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcloudmanagementsdkgo.New()

    res, err := s.Workspaces.Create(ctx, operations.V1WorkspacesCreateWorkspaceSecurity{
        Oauth2: &components.SchemeOauth2{
            ClientID: os.Getenv("CRIBLMGMTPLANE_CLIENT_ID"),
            ClientSecret: os.Getenv("CRIBLMGMTPLANE_CLIENT_SECRET"),
            TokenURL: os.Getenv("CRIBLMGMTPLANE_TOKEN_URL"),
            Audience: "https://publicapi.cribl.cloud",
        },
    }, "<id>", components.WorkspaceCreateRequestDTO{
        WorkspaceID: "main",
        Region: components.WorkspaceCreateRequestDTORegionUsWest2,
        Alias: criblcloudmanagementsdkgo.String("Production Environment"),
        Description: criblcloudmanagementsdkgo.String("Main production workspace for customer data processing"),
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `security`                                                                                                       | [operations.V1WorkspacesCreateWorkspaceSecurity](../../models/operations/v1workspacescreateworkspacesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `organizationID`                                                                                                 | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Organization identifier                                                                                          |
| `workspaceCreateRequestDTO`                                                                                      | [components.WorkspaceCreateRequestDTO](../../models/components/workspacecreaterequestdto.md)                     | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

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
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcloudmanagementsdkgo.New()

    res, err := s.Workspaces.List(ctx, operations.V1WorkspacesListWorkspacesSecurity{
        Oauth2: &components.SchemeOauth2{
            ClientID: os.Getenv("CRIBLMGMTPLANE_CLIENT_ID"),
            ClientSecret: os.Getenv("CRIBLMGMTPLANE_CLIENT_SECRET"),
            TokenURL: os.Getenv("CRIBLMGMTPLANE_TOKEN_URL"),
            Audience: "https://publicapi.cribl.cloud",
        },
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspacesListResponseDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `security`                                                                                                     | [operations.V1WorkspacesListWorkspacesSecurity](../../models/operations/v1workspaceslistworkspacessecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |
| `organizationID`                                                                                               | *string*                                                                                                       | :heavy_check_mark:                                                                                             | Organization identifier                                                                                        |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

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
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcloudmanagementsdkgo.New()

    res, err := s.Workspaces.Update(ctx, operations.V1WorkspacesUpdateWorkspaceSecurity{
        Oauth2: &components.SchemeOauth2{
            ClientID: os.Getenv("CRIBLMGMTPLANE_CLIENT_ID"),
            ClientSecret: os.Getenv("CRIBLMGMTPLANE_CLIENT_SECRET"),
            TokenURL: os.Getenv("CRIBLMGMTPLANE_TOKEN_URL"),
            Audience: "https://publicapi.cribl.cloud",
        },
    }, "<id>", "<id>", components.WorkspacePatchRequestDTO{
        Alias: criblcloudmanagementsdkgo.String("Production Environment"),
        Description: criblcloudmanagementsdkgo.String("Main production workspace for customer data processing"),
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `security`                                                                                                       | [operations.V1WorkspacesUpdateWorkspaceSecurity](../../models/operations/v1workspacesupdateworkspacesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `organizationID`                                                                                                 | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Organization identifier                                                                                          |
| `workspaceID`                                                                                                    | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Workspace identifier                                                                                             |
| `workspacePatchRequestDTO`                                                                                       | [components.WorkspacePatchRequestDTO](../../models/components/workspacepatchrequestdto.md)                       | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

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
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcloudmanagementsdkgo.New()

    res, err := s.Workspaces.Delete(ctx, operations.V1WorkspacesDeleteWorkspaceSecurity{
        Oauth2: &components.SchemeOauth2{
            ClientID: os.Getenv("CRIBLMGMTPLANE_CLIENT_ID"),
            ClientSecret: os.Getenv("CRIBLMGMTPLANE_CLIENT_SECRET"),
            TokenURL: os.Getenv("CRIBLMGMTPLANE_TOKEN_URL"),
            Audience: "https://publicapi.cribl.cloud",
        },
    }, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `security`                                                                                                       | [operations.V1WorkspacesDeleteWorkspaceSecurity](../../models/operations/v1workspacesdeleteworkspacesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `organizationID`                                                                                                 | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Organization identifier                                                                                          |
| `workspaceID`                                                                                                    | *string*                                                                                                         | :heavy_check_mark:                                                                                               | Workspace identifier                                                                                             |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

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
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"os"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcloudmanagementsdkgo.New()

    res, err := s.Workspaces.Get(ctx, operations.V1WorkspacesGetWorkspaceSecurity{
        Oauth2: &components.SchemeOauth2{
            ClientID: os.Getenv("CRIBLMGMTPLANE_CLIENT_ID"),
            ClientSecret: os.Getenv("CRIBLMGMTPLANE_CLIENT_SECRET"),
            TokenURL: os.Getenv("CRIBLMGMTPLANE_TOKEN_URL"),
            Audience: "https://publicapi.cribl.cloud",
        },
    }, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `security`                                                                                                 | [operations.V1WorkspacesGetWorkspaceSecurity](../../models/operations/v1workspacesgetworkspacesecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `organizationID`                                                                                           | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Organization identifier                                                                                    |
| `workspaceID`                                                                                              | *string*                                                                                                   | :heavy_check_mark:                                                                                         | Workspace identifier                                                                                       |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.V1WorkspacesGetWorkspaceResponse](../../models/operations/v1workspacesgetworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |