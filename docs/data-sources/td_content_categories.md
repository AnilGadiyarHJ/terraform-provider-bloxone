---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_td_content_categories Data Source - terraform-provider-bloxone"
subcategory: "Threat Defense"
description: |-
  Retrieves information about existing Content Categories.
---

# bloxone_td_content_categories (Data Source)

Retrieves information about existing Content Categories.

## Example Usage

```terraform
# Get all Content Categories
data "bloxone_td_content_categories" "all_content_categories" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Optional:

- `category_code` (Number) The category code.
- `category_name` (String) The name of the category.
- `functional_group` (String) The functional group name of the category.