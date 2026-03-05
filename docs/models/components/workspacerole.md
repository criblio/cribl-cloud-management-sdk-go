# WorkspaceRole

Role assigned to the API Credential on the Workspace.

## Example Usage

```go
import (
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
)

value := components.WorkspaceRoleOwner

// Open enum: custom values can be created with a direct type cast
custom := components.WorkspaceRole("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `WorkspaceRoleOwner`    | owner                   |
| `WorkspaceRoleAdmin`    | admin                   |
| `WorkspaceRoleUser`     | user                    |
| `WorkspaceRoleNoaccess` | noaccess                |