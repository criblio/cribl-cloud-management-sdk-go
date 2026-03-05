# Region

AWS region where the Workspace is deployed.

## Example Usage

```go
import (
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
)

value := components.RegionUsWest2

// Open enum: custom values can be created with a direct type cast
custom := components.Region("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `RegionUsWest2`      | us-west-2            |
| `RegionUsEast1`      | us-east-1            |
| `RegionUsEast2`      | us-east-2            |
| `RegionEuCentral1`   | eu-central-1         |
| `RegionEuCentral2`   | eu-central-2         |
| `RegionEuWest2`      | eu-west-2            |
| `RegionApSoutheast1` | ap-southeast-1       |
| `RegionApSoutheast2` | ap-southeast-2       |
| `RegionCaCentral1`   | ca-central-1         |
| `RegionApNortheast1` | ap-northeast-1       |
| `RegionSaEast1`      | sa-east-1            |
| `RegionEuWest1`      | eu-west-1            |
| `RegionEuWest3`      | eu-west-3            |