---
layout: ""
page_title: dynatrace_oneagent_default_version Resource - terraform-provider-dynatrace"
subcategory: "Environment Settings"
description: |-
  The resource `dynatrace_oneagent_default_version` covers configuration for OneAgent default version
---

# dynatrace_oneagent_default_version (Resource)

-> This resource requires the API token scopes **Read settings** (`settings.read`) and **Write settings** (`settings.write`)

## Dynatrace Documentation

- OneAgent update - https://www.dynatrace.com/support/help/setup-and-configuration/dynatrace-oneagent/oneagent-update

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:deployment.oneagent.default-version`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_oneagent_default_version` downloads existing OneAgent default version configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_oneagent_default_version" "#name#" {
  default_version = "latest"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `default_version` (String) Default version

### Optional

- `revision` (String) Revision

### Read-Only

- `id` (String) The ID of this resource.
 