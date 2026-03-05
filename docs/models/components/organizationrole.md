# OrganizationRole

Organization-level Role assigned to the API Credential.

## Example Usage

```go
import (
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
)

value := components.OrganizationRoleOwner

// Open enum: custom values can be created with a direct type cast
custom := components.OrganizationRole("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `OrganizationRoleOwner` | owner                   |
| `OrganizationRoleAdmin` | admin                   |
| `OrganizationRoleUser`  | user                    |