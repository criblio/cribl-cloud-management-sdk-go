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
			ClientOauth: &components.SchemeClientOauth{
				ClientID:     os.Getenv("CRIBLMGMTPLANE_CLIENT_ID"),
				ClientSecret: os.Getenv("CRIBLMGMTPLANE_CLIENT_SECRET"),
				TokenURL:     os.Getenv("CRIBLMGMTPLANE_TOKEN_URL"),
				Audience:     "https://api.cribl.cloud",
			},
		}),
	)

	res, err := s.Health.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->