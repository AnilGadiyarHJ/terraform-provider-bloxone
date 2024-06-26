---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dhcp_hosts Data Source - terraform-provider-bloxone"
subcategory: "IPAM/DHCP"
description: |-
  Retrieves information about existing DHCP Hosts.
  A DHCP Host object associates a DHCP Config Profile with an on-prem host.
---

# bloxone_dhcp_hosts (Data Source)

Retrieves information about existing DHCP Hosts.

A DHCP Host object associates a DHCP Config Profile with an on-prem host.

## Example Usage

```terraform
# Get DHCP Host filtered by an attribute
data "bloxone_dhcp_hosts" "example_by_name" {
  filters = {
    name = "example-host"
  }
}

# Get DHCP Host by tag
data "bloxone_dhcp_hosts" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all DHCP hosts
data "bloxone_dhcp_hosts" "example_all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filters` (Map of String) Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `retry_if_not_found` (Boolean) If set to `true`, the data source will retry until a matching host is found, or until the Read Timeout expires.
- `tag_filters` (Map of String) Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

### Read-Only

- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String) [Duration](https://pkg.go.dev/time#ParseDuration) to wait before being considered a timeout during read operations. Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Default is 20m.


<a id="nestedatt--results"></a>
### Nested Schema for `results`

Optional:

- `associated_server` (Attributes) The DHCP Config Profile for the on-prem host. (see [below for nested schema](#nestedatt--results--associated_server))
- `ip_space` (String) The resource identifier.
- `server` (String) The resource identifier.
- `tags` (Map of String) The tags of the on-prem host in JSON format.

Read-Only:

- `address` (String) The primary IP address of the on-prem host.
- `anycast_addresses` (List of String) Anycast address configured to the host. Order is not significant.
- `comment` (String) The description for the on-prem host.
- `current_version` (String) Current dhcp application version of the host.
- `id` (String) The resource identifier.
- `name` (String) The display name of the on-prem host.
- `ophid` (String) The on-prem host ID.
- `provider_id` (String) External provider identifier.
- `tags_all` (Map of String) The tags of the on-prem host in JSON format including default tags.
- `type` (String) Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory.

<a id="nestedatt--results--associated_server"></a>
### Nested Schema for `results.associated_server`

Read-Only:

- `id` (String) The resource identifier.
- `name` (String) The DHCP Config Profile name.
