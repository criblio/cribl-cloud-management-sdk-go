# CreditGrantSchema


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Date`                                             | [time.Time](https://pkg.go.dev/time#Time)          | :heavy_check_mark:                                 | Date the grant was issued, in ISO 8601 format.     | 2025-10-29 00:00:00 +0000 UTC                      |
| `CreditsAcquired`                                  | `float64`                                          | :heavy_check_mark:                                 | Number of credits acquired by the grant.           | 6000000                                            |
| `Type`                                             | [components.Type](../../models/components/type.md) | :heavy_check_mark:                                 | Type of the credit grant.                          | Purchased                                          |