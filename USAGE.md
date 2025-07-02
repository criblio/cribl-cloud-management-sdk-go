<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := criblcloudmanagementsdkgo.New(
		"https://api.example.com",
		criblcloudmanagementsdkgo.WithSecurity(os.Getenv("CRIBLMGMTPLANE_BEARER_AUTH")),
	)

	res, err := s.DummyServiceStatus(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->