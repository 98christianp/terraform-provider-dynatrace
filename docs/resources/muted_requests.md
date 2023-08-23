---
layout: ""
page_title: dynatrace_muted_requests Resource - terraform-provider-dynatrace"
subcategory: "Service Monitoring"
description: |-
  The resource `dynatrace_muted_requests` covers configuration for muted requests
---

# dynatrace_muted_requests (Resource)

-> This resource requires the API token scopes **Read settings** (`settings.read`) and **Write settings** (`settings.write`)

## Dynatrace Documentation

- Mute monitoring of service requests - https://www.dynatrace.com/support/help/how-to-use-dynatrace/services/service-monitoring-settings/service-monitoring-mute 

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:settings.mutedrequests`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_muted_requests` downloads all existing muted requests configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_muted_requests" "#name#" {
  muted_request_names = [ "/healthcheck", "/heartbeat" ]
  service_id          = "SERVICE-1234567890000000"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `service_id` (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.

### Optional

- `muted_request_names` (Set of String) Muted request names

### Read-Only

- `id` (String) The ID of this resource.
 