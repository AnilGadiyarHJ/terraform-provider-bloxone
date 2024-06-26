---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_ipam_hosts Data Source - terraform-provider-bloxone"
subcategory: "IPAM/DHCP"
description: |-
  Retrieves information about existing IPAM Hosts.
  The IPAM host object represents any network connected equipment that is assigned one or more IP Addresses.
---

# bloxone_ipam_hosts (Data Source)

Retrieves information about existing IPAM Hosts.

The IPAM host object represents any network connected equipment that is assigned one or more IP Addresses.

## Example Usage

```terraform
# Get IPAM Hosts filtered by an attribute
data "bloxone_ipam_hosts" "example_by_attribute" {
  filters = {
    "name" = "example_ipam_host"
  }
}

# Get IPAM Hosts filtered by tag
data "bloxone_ipam_hosts" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all IPAM Hosts
data "bloxone_ipam_hosts" "example_all" {}
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

- `name` (String) The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8.

Optional:

- `addresses` (Attributes List) The list of all addresses associated with the IPAM host, which may be in different IP spaces. (see [below for nested schema](#nestedatt--results--addresses))
- `auto_generate_records` (Boolean) This flag specifies if resource records have to be auto generated for the host.
- `comment` (String) The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8.
- `host_names` (Attributes List) The name records to be generated for the host.  This field is required if _auto_generate_records_ is true. (see [below for nested schema](#nestedatt--results--host_names))
- `tags` (Map of String) The tags for the IPAM host in JSON format.

Read-Only:

- `created_at` (String) Time when the object has been created.
- `id` (String) The resource identifier.
- `tags_all` (Map of String) The tags for the IPAM host in JSON format including default tags.
- `updated_at` (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.

<a id="nestedatt--results--addresses"></a>
### Nested Schema for `results.addresses`

Optional:

- `address` (String) Field usage depends on the operation:
  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.
  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:
    * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.
    * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.
    * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.
- `next_available_id` (String) The resource identifier for the network container where the next available address should be generated for the host
- `ref` (String) The resource identifier.
- `space` (String) The resource identifier.


<a id="nestedatt--results--host_names"></a>
### Nested Schema for `results.host_names`

Required:

- `name` (String) A name for the host.
- `zone` (String) The resource identifier.

Optional:

- `alias` (Boolean) When _true_, the name is treated as an alias.
- `primary_name` (Boolean) When _true_, the name field is treated as primary name. There must be one and only one primary name in the list of host names. The primary name will be treated as the canonical name for all the aliases. PTR record will be generated only for the primary name.
