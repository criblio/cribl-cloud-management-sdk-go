/*
 * Cribl.Cloud Authentication Example
 *
 * This example demonstrates how to configure authentication on Cribl.Cloud
 * using OAuth2 credentials.
 *
 * 1. Create an SDK client with OAuth2 client credentials using the
 * client_oauth security scheme.
 * 2. Automatically handle token exchange and refresh.
 * 3. Validate the connection by listing Workspaces.
 *
 * Prerequisites: Replace the placeholder values for ORG_ID, CLIENT_ID, and
 * CLIENT_SECRET with your Organization ID and Client ID and Secret.
 * To get your Client ID and Secret, follow the steps at
 * https://docs.cribl.io/cribl-as-code/sdks-auth/#sdks-auth-cloud.
 * Your Client ID and Secret are sensitive information and should be kept private.
 *
 * NOTE: This example is for Cribl.Cloud only.
 * It does not require .env file configuration.
 */

package main

import (
	"context"
	"fmt"
	"os"

	criblcloudmanagementsdkgo "github.com/criblio/cribl-cloud-management-sdk-go"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/apierrors"
	"github.com/criblio/cribl-cloud-management-sdk-go/models/components"
)

// Cribl.Cloud configuration: Replace the placeholder values
const (
	CLIENT_ID     = "your-client-id"     // Replace with your OAuth2 Client ID
	CLIENT_SECRET = "your-client-secret" // Replace with your OAuth2 Client Secret
	ORG_ID        = "your-org-id"        // Replace with your Organization ID

	// Token URL and audience for Cribl Cloud OAuth2
	TOKEN_URL = "https://login.cribl.cloud/oauth/token"
	AUDIENCE  = "https://api.cribl.cloud"
)

func main() {
	// Create authenticated SDK client with OAuth2
	criblMgmtPlane := criblcloudmanagementsdkgo.New(
		criblcloudmanagementsdkgo.WithSecurity(components.Security{
			ClientOauth: &components.SchemeClientOauth{
				ClientID:     CLIENT_ID,
				ClientSecret: CLIENT_SECRET,
				TokenURL:     TOKEN_URL,
				Audience:     AUDIENCE,
			},
		}),
	)

	fmt.Println("✅ Cribl.Cloud Management Plane SDK client created")

	// Create context for API calls
	ctx := context.Background()

	// List all Workspaces for the Organization
	workspacesResponse, err := criblMgmtPlane.Workspaces.List(ctx, ORG_ID)
	if err != nil {
		handleError(err)
		return
	}

	// Check if response is successful
	if workspacesResponse.WorkspacesListResponseDTO != nil {
		items := workspacesResponse.WorkspacesListResponseDTO.Items
		count := workspacesResponse.WorkspacesListResponseDTO.Count

		var message string
		if count > 0 {
			message = fmt.Sprintf("✅ Client works! Found %.0f Workspace(s):", count)
		} else {
			message = fmt.Sprintf("✅ Client works! No Workspaces found for Organization %s", ORG_ID)
		}
		fmt.Println(message)

		for _, ws := range items {
			aliasDisplay := ""
			if ws.Alias != nil {
				aliasDisplay = fmt.Sprintf(" (%s)", *ws.Alias)
			}
			fmt.Printf("  • %s%s - %s in %s\n", ws.WorkspaceID, aliasDisplay, ws.State, ws.Region)
		}
	} else if workspacesResponse.DefaultErrorDTO != nil {
		fmt.Printf("❌ Error listing Workspaces: %v\n", workspacesResponse.DefaultErrorDTO)
		os.Exit(1)
	}
}

func handleError(err error) {
	// Check if it's an API error
	if apiErr, ok := err.(*apierrors.APIError); ok {
		statusCode := apiErr.StatusCode
		if statusCode == 401 {
			fmt.Println("⚠️ Authentication failed! Check your CLIENT_ID and CLIENT_SECRET.")
		} else if statusCode == 429 {
			fmt.Println("⚠️ Uh oh, you've reached the rate limit! Try again in a few seconds.")
		} else {
			fmt.Printf("❌ Something went wrong: %v\n", err)
		}
	} else {
		fmt.Printf("❌ Something went wrong: %v\n", err)
	}
	os.Exit(1)
}