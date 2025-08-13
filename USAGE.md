<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcloudmanagementsdkgo.New(
		criblcloudmanagementsdkgo.WithSecurity(components.Security{
			ClientID:     criblcloudmanagementsdkgo.String(os.Getenv("CRIBLMGMTPLANE_CLIENT_ID")),
			ClientSecret: criblcloudmanagementsdkgo.String(os.Getenv("CRIBLMGMTPLANE_CLIENT_SECRET")),
			Audience:     criblcloudmanagementsdkgo.String("https://publicapi.cribl.cloud"),
		}),
	)

	res, err := s.Health.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->