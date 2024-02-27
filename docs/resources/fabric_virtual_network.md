---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_fabric_virtual_network Resource - terraform-provider-catalystcenter"
subcategory: "SDA"
description: |-
  This resource can manage a Fabric Virtual Network.
---

# catalystcenter_fabric_virtual_network (Resource)

This resource can manage a Fabric Virtual Network.

## Example Usage

```terraform
resource "catalystcenter_fabric_virtual_network" "example" {
  virtual_network_name = "SDA_VN1"
  is_guest             = false
  sg_names             = ["Employees"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `virtual_network_name` (String) Virtual Network Name to be assigned at global level

### Optional

- `is_guest` (Boolean) Guest Virtual Network enablement flag
  - Default value: `false`
- `sg_names` (List of String) Scalable Groups to be associated to virtual network
- `vmanage_vpn_id` (String) vManage vpn id for SD-WAN

### Read-Only

- `id` (String) The id of the object

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_fabric_virtual_network.example "SDA_VN1"
```