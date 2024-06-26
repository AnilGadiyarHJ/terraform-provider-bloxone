---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_dns_delegation Resource - terraform-provider-bloxone"
subcategory: "DNS"
description: |-
  Manages a zone delegation.
---

# bloxone_dns_delegation (Resource)

Manages a zone delegation.

## Example Usage

```terraform
resource "bloxone_dns_view" "example" {
  name = "example-view"
}

resource "bloxone_dns_auth_zone" "example" {
  fqdn         = "domain.com."
  primary_type = "cloud"
  view         = bloxone_dns_view.example.id
}

resource "bloxone_dns_delegation" "example" {
  fqdn = "del.domain.com."
  delegation_servers = [{
    address = "12.0.0.0"
    fqdn    = "ns1.com."
  }]

  # Other optional fields
  view    = bloxone_dns_view.example.id
  comment = "Delegation zone created through Terraform"
  tags = {
    site = "Site A"
  }
  disabled = true

  depends_on = [bloxone_dns_view.example, bloxone_dns_auth_zone.example]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `delegation_servers` (Attributes List) Required. DNS zone delegation servers. Order is not significant. (see [below for nested schema](#nestedatt--delegation_servers))
- `fqdn` (String) Delegation FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.

### Optional

- `comment` (String) Optional. Comment for zone delegation.
- `disabled` (Boolean) Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records.
- `tags` (Map of String) Tagging specifics.
- `view` (String) The resource identifier.

### Read-Only

- `id` (String) The resource identifier.
- `parent` (String) The resource identifier.
- `protocol_fqdn` (String) Delegation FQDN in punycode.
- `tags_all` (Map of String) Tagging specifics includes default tags.

<a id="nestedatt--delegation_servers"></a>
### Nested Schema for `delegation_servers`

Required:

- `fqdn` (String) Required. FQDN of nameserver.

Optional:

- `address` (String) Optional. IP Address of nameserver.  Only required when fqdn of a delegation server falls under delegation fqdn

Read-Only:

- `protocol_fqdn` (String) FQDN of nameserver in punycode.
