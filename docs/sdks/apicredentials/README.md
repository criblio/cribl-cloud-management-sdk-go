# ApiCredentials

## Overview

Operations related to API credentials

### Available Operations

* [List](#list) - List API Credentials for an Organization
* [Create](#create) - Create an API Credential
* [Update](#update) - Update an API Credential
* [Delete](#delete) - Delete an API Credential
* [Get](#get) - Get an API Credential

## List

Get a list of all API Credentials for the specified Organization.

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

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |
| `organizationID`                                                                | *string*                                                                        | :heavy_check_mark:                                                              | The <code>id</code> of the Organization whose API Credentials you want to list. |
| `opts`                                                                          | [][operations.Option](../../models/operations/option.md)                        | :heavy_minus_sign:                                                              | The options for this request.                                                   |

### Response

**[*operations.V1APICredentialsListAPICredentialsResponse](../../models/operations/v1apicredentialslistapicredentialsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new API Credential for the specified Organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.apiCredentials.createApiCredential" method="post" path="/v1/organizations/{organizationId}/api-credentials" -->
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

    res, err := s.APICredentials.Create(ctx, "<id>", components.APICredentialCreateRequestDTO{
        Name: "Auto-Manage-Workspaces",
        Description: "Used for automated Workspace management",
        Enabled: true,
        WorkspaceID: "main",
        Roles: components.APICredentialRolesSchema{
            OrganizationRole: components.OrganizationRoleAdmin,
            Workspaces: []components.WorkspaceRoleSchema{
                components.WorkspaceRoleSchema{
                    WorkspaceID: "main",
                    WorkspaceRole: components.WorkspaceRoleAdmin,
                    Products: []components.ProductRoleSchema{
                        components.ProductRoleSchema{
                            Product: components.ProductStream,
                            Role: components.RoleAdmin,
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APICredentialResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `organizationID`                                                                                     | *string*                                                                                             | :heavy_check_mark:                                                                                   | The <code>id</code> of the Organization where you want to create the API Credential.                 |
| `apiCredentialCreateRequestDTO`                                                                      | [components.APICredentialCreateRequestDTO](../../models/components/apicredentialcreaterequestdto.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V1APICredentialsCreateAPICredentialResponse](../../models/operations/v1apicredentialscreateapicredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified API Credential.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.apiCredentials.updateApiCredential" method="patch" path="/v1/organizations/{organizationId}/api-credentials/{apiCredentialId}" -->
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

    res, err := s.APICredentials.Update(ctx, "<id>", "<id>", components.APICredentialUpdateRequestDTO{
        Name: criblcloudmanagementsdkgo.Pointer("Auto-Manage-Workspaces"),
        Description: criblcloudmanagementsdkgo.Pointer("Used for automated Workspace management"),
        Enabled: criblcloudmanagementsdkgo.Pointer(true),
        Roles: &components.APICredentialRolesSchema{
            OrganizationRole: components.OrganizationRoleAdmin,
            Workspaces: []components.WorkspaceRoleSchema{
                components.WorkspaceRoleSchema{
                    WorkspaceID: "main",
                    WorkspaceRole: components.WorkspaceRoleAdmin,
                    Products: []components.ProductRoleSchema{
                        components.ProductRoleSchema{
                            Product: components.ProductStream,
                            Role: components.RoleAdmin,
                        },
                    },
                },
            },
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `organizationID`                                                                                     | *string*                                                                                             | :heavy_check_mark:                                                                                   | The <code>id</code> of the Organization whose API Credential you want to update.                     |
| `apiCredentialID`                                                                                    | *string*                                                                                             | :heavy_check_mark:                                                                                   | The <code>clientId</code> of the API Credential to update.                                           |
| `apiCredentialUpdateRequestDTO`                                                                      | [components.APICredentialUpdateRequestDTO](../../models/components/apicredentialupdaterequestdto.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.V1APICredentialsUpdateAPICredentialResponse](../../models/operations/v1apicredentialsupdateapicredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified API Credential.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.apiCredentials.deleteApiCredential" method="delete" path="/v1/organizations/{organizationId}/api-credentials/{apiCredentialId}" -->
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

    res, err := s.APICredentials.Delete(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.DefaultErrorDTO != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `organizationID`                                                                 | *string*                                                                         | :heavy_check_mark:                                                               | The <code>id</code> of the Organization whose API Credential you want to delete. |
| `apiCredentialID`                                                                | *string*                                                                         | :heavy_check_mark:                                                               | The <code>clientId</code> of the API Credential to delete.                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.V1APICredentialsDeleteAPICredentialResponse](../../models/operations/v1apicredentialsdeleteapicredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified API Credential.

### Example Usage

<!-- UsageSnippet language="go" operationID="v1.apiCredentials.getApiCredential" method="get" path="/v1/organizations/{organizationId}/api-credentials/{apiCredentialId}" -->
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

    res, err := s.APICredentials.Get(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.APICredentialResponseSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `organizationID`                                                                   | *string*                                                                           | :heavy_check_mark:                                                                 | The <code>id</code> of the Organization whose API Credential you want to retrieve. |
| `apiCredentialID`                                                                  | *string*                                                                           | :heavy_check_mark:                                                                 | The <code>clientId</code> of the API Credential to retrieve.                       |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V1APICredentialsGetAPICredentialResponse](../../models/operations/v1apicredentialsgetapicredentialresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |