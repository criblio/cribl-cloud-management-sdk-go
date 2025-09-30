# WorkspaceCreateRequestDTO


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `WorkspaceID`                                          | *string*                                               | :heavy_check_mark:                                     | Unique identifier for the workspace                    | main                                                   |
| `Alias`                                                | **string*                                              | :heavy_minus_sign:                                     | User-friendly alias for the workspace                  | Production Environment                                 |
| `Description`                                          | **string*                                              | :heavy_minus_sign:                                     | Detailed description of the workspace                  | Main production workspace for customer data processing |
| `Tags`                                                 | []*string*                                             | :heavy_minus_sign:                                     | Tags associated with the workspace                     | [<br/>"production",<br/>"customer-data"<br/>]          |