# State

Current state of the Workspace.

## Example Usage

```go
import (
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
)

value := components.StateProvisioning

// Open enum: custom values can be created with a direct type cast
custom := components.State("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `StateProvisioning`   | Provisioning          |
| `StateActive`         | Active                |
| `StateInactive`       | Inactive              |
| `StateFailed`         | Failed                |
| `StateDeprovisioning` | Deprovisioning        |