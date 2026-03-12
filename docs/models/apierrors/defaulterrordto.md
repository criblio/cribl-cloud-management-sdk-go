# DefaultErrorDTO


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `StatusCode`                                                       | `int64`                                                            | :heavy_check_mark:                                                 | HTTP status code.                                                  | 400                                                                |
| `Message`                                                          | `string`                                                           | :heavy_check_mark:                                                 | Human-readable error message.                                      | Bad Request                                                        |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |