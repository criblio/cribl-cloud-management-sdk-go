# WorkspaceCreateRequestDTO


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `WorkspaceID`                                          | *string*                                               | :heavy_check_mark:                                     | Unique identifier for the Workspace.                   | main                                                   |
| `Alias`                                                | **string*                                              | :heavy_minus_sign:                                     | User-friendly alias for the Workspace.                 | Production Environment                                 |
| `Description`                                          | **string*                                              | :heavy_minus_sign:                                     | Brief description of the Workspace.                    | Main production Workspace for customer data processing |
| `Tags`                                                 | []*string*                                             | :heavy_minus_sign:                                     | Tags associated with the Workspace.                    | [<br/>"production",<br/>"customer-data"<br/>]          |