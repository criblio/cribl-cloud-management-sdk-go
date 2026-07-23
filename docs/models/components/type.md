# Type

Type of the credit grant.

## Example Usage

```go
import (
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
)

value := components.TypeRefund

// Open enum: custom values can be created with a direct type cast
custom := components.Type("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `TypeRefund`    | Refund          |
| `TypeRollover`  | Rollover        |
| `TypePurchased` | Purchased       |