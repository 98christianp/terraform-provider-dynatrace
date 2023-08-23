---
layout: ""
page_title: "dynatrace_slo_normalization Resource - terraform-provider-dynatrace"
subcategory: "Service-level Objective"
description: |-
  The resource `dynatrace_slo_normalization` covers configuration for service-level objective setup
---

# dynatrace_slo_normalization (Resource)

-> This resource requires the API token scopes **Read settings** (`settings.read`) and **Write settings** (`settings.write`)

## Dynatrace Documentation

- Normalize error budget - https://www.dynatrace.com/support/help/platform-modules/cloud-automation/service-level-objectives/configure-and-monitor-slo#normalize-error-budget

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:monitoring.slo.normalization`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_slo_normalization` downloads all existing service-level objective setup configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_slo_normalization" "#name#" {
  normalize = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `normalize` (Boolean) When set to true, the error budget left will be shown in percent of the total error budget. For more details see [SLO normalization help](https://dt-url.net/slo-normalize-error-budget).

### Read-Only

- `id` (String) The ID of this resource.
 