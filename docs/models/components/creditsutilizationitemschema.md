# CreditsUtilizationItemSchema


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Date`                                                | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | Date of the aggregation window in ISO 8601 format.    | 2025-05-01 00:00:00 +0000 UTC                         |
| `Credits`                                             | `float64`                                             | :heavy_check_mark:                                    | Number of credits consumed in the aggregation window. | 147722.91698535177                                    |