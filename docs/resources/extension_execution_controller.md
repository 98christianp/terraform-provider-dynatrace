---
layout: ""
page_title: dynatrace_extension_execution_controller Resource - terraform-provider-dynatrace"
subcategory: "Extensions"
description: |-
  The resource `dynatrace_extension_execution_controller` covers Extension Execution Controller configuration for OneAgent deployment
---

# dynatrace_extension_execution_controller (Resource)

-> This resource requires the API token scopes **Read settings** (`settings.read`) and **Write settings** (`settings.write`)

## Dynatrace Documentation

- Extensions 2.0 - https://www.dynatrace.com/support/help/extend-dynatrace/extensions20/extensions-concepts

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:eec.local`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_extension_execution_controller` downloads all existing Extension Execution Controller configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_extension_execution_controller" "#name#" {
  enabled             = true
  ingest_active       = false
  performance_profile = "DEFAULT"
  scope               = "environment"
  statsd_active       = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)

### Optional

- `ingest_active` (Boolean) Enable local HTTP Metric, Log and Event Ingest API
- `performance_profile` (String) Possible Values: `DEFAULT`, `HIGH`
- `scope` (String) The scope of this setting (HOST, HOST_GROUP). Omit this property if you want to cover the whole environment.
- `statsd_active` (Boolean) Enable Dynatrace StatsD

### Read-Only

- `id` (String) The ID of this resource.
 