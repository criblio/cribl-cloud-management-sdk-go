# cribl-cloud-management-sdk-go Examples

This directory contains example programs demonstrating how to use the Cribl Cloud Management SDK for Go.

## Prerequisites

- Go 1.22 or higher
- A Cribl Cloud account with API credentials

## Available Example

### Authentication Example (`auth-example.go`)

This example demonstrates how to configure authentication on Cribl.Cloud using OAuth2 credentials. It performs the following steps:

1. Create an SDK client with OAuth2 client credentials using the `client_oauth` security scheme
2. Automatically handle token exchange and refresh
3. Validate the connection by listing Workspaces

**Note**: This example is for Cribl.Cloud only and does not require `.env` file configuration.

## Setup

Before running the example, you need to obtain your Cribl.Cloud API credentials:

1. **Get your Organization ID**: Available in your Cribl.Cloud console
2. **Create API credentials**: Follow the steps at [https://docs.cribl.io/cribl-as-code/sdks-auth/#sdks-auth-cloud](https://docs.cribl.io/cribl-as-code/sdks-auth/#sdks-auth-cloud) to get your Client ID and Secret

> **⚠️ Important**: Your Client ID and Secret are sensitive information and should be kept private.

## Configuration

Open `auth-example.go` and replace the placeholder values with your actual credentials:

```go
const (
	CLIENT_ID     = "your-client-id"     // Replace with your OAuth2 Client ID
	CLIENT_SECRET = "your-client-secret" // Replace with your OAuth2 Client Secret
	ORG_ID        = "your-org-id"        // Replace with your Organization ID

	// Token URL and audience for Cribl Cloud OAuth2 (predefined)
	TOKEN_URL = "https://login.cribl.cloud/oauth/token"
	AUDIENCE  = "https://api.cribl.cloud"
)
```

> **⚠️ Security Note**: Never commit your actual credentials to version control. Consider using environment variables in production applications.

## Running the Example

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the authentication example:
   ```bash
   go run auth-example.go
   ```

## Expected Output

On successful authentication, you should see output similar to:

```
✅ Cribl.Cloud Management Plane SDK client created
✅ Client works! Found 2 Workspace(s):
  • ws-12345 (My Workspace) - running in us-west-2
  • ws-67890 (Test Environment) - stopped in us-east-1
```

## Troubleshooting

- **Authentication failed**: Verify your `CLIENT_ID`, `CLIENT_SECRET`, and `ORG_ID` are correct
- **Rate limit reached**: Wait a few seconds and try again
- **Network errors**: Check your internet connection and firewall settings

## Building the Example

To build a standalone executable:

```bash
go build -o auth-example auth-example.go
```

## Next Steps

After successfully running this authentication example, you can:
- Explore the available SDK methods in the [main project documentation](../README.md)
- Check out the [API reference](https://docs.cribl.io/cribl-as-code/api-reference) for detailed information about available endpoints
- Build your own applications using the authenticated SDK client pattern shown in this example
