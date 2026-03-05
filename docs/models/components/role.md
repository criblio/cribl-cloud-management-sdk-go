# Role

Role assigned to the API Credential on the product.

## Example Usage

```go
import (
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
)

value := components.RoleAdmin

// Open enum: custom values can be created with a direct type cast
custom := components.Role("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `RoleAdmin`    | admin          |
| `RoleEditor`   | editor         |
| `RoleReader`   | reader         |
| `RoleUser`     | user           |
| `RoleNoaccess` | noaccess       |