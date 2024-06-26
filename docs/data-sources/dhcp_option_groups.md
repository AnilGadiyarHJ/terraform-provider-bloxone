---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dhcp_option_groups Data Source - terraform-provider-bloxone"
subcategory: "IPAM/DHCP"
description: |-
  Retrieves information about existing DHCP Option Groups.
  The Option Group object is a named collection of options.
---

# bloxone_dhcp_option_groups (Data Source)

Retrieves information about existing DHCP Option Groups.

The Option Group object is a named collection of options.

## Example Usage

```terraform
# Get DHCP Option groups filtered by an attribute
data "bloxone_dhcp_option_groups" "example_by_name" {
  filters = {
    name = "example-group"
  }
}

# Get DHCP Option group/s by tag
data "bloxone_dhcp_option_groups" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all DHCP Option groups
data "bloxone_dhcp_option_groups" "example_all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filters` (Map of String) Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `tag_filters` (Map of String) Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.

### Read-Only

- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Required:

- `name` (String) The name of the option group. Must contain 1 to 256 characters. Can include UTF-8.
- `protocol` (String) The type of protocol (_ip4_ or _ip6_).

Optional:

- `comment` (String) The description for the option group. May contain 0 to 1024 characters. Can include UTF-8.
- `dhcp_options` (Attributes List) The list of DHCP options for the option group. May be either a specific option or a group of options. (see [below for nested schema](#nestedatt--results--dhcp_options))
- `tags` (Map of String) The tags for the option group in JSON format.

Read-Only:

- `created_at` (String) Time when the object has been created.
- `id` (String) The resource identifier.
- `tags_all` (Map of String) The tags for the option group in JSON format including default tag.
- `updated_at` (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.

<a id="nestedatt--results--dhcp_options"></a>
### Nested Schema for `results.dhcp_options`

Optional:

- `group` (String) The resource identifier.
- `option_code` (String) The resource identifier.
- `option_value` (String) The option value.
- `type` (String) The type of item. Valid values are:
  * _group_
  * _option_
