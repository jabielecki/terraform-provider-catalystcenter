---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_deploy_template Resource - terraform-provider-catalystcenter"
subcategory: "Templates"
description: |-
  This resource can manage a Deploy Template.
---

# catalystcenter_deploy_template (Resource)

This resource can manage a Deploy Template.

## Example Usage

```terraform
resource "catalystcenter_deploy_template" "example" {
  template_id         = "12345678-1234-1234-1234-123456789012"
  force_push_template = false
  is_composite        = false
  target_info = [
    {
      host_name             = "SW01"
      type                  = "MANAGED_DEVICE_HOSTNAME"
      versioned_template_id = "12345678-1234-1234-1234-123456789012"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `target_info` (Attributes List) Target info to deploy template (see [below for nested schema](#nestedatt--target_info))
- `template_id` (String) ID of template to be provisioned

### Optional

- `force_push_template` (Boolean) Force Push Template
- `is_composite` (Boolean) Composite template flag
- `main_template_id` (String) Composite Template ID
- `member_template_deployment_info` (Attributes List) Member Template Deployment Info (see [below for nested schema](#nestedatt--member_template_deployment_info))

### Read-Only

- `id` (String) The id of the object

<a id="nestedatt--target_info"></a>
### Nested Schema for `target_info`

Required:

- `type` (String) Target type of device
  - Choices: `MANAGED_DEVICE_IP`, `MANAGED_DEVICE_UUID`, `PRE_PROVISIONED_SERIAL`, `PRE_PROVISIONED_MAC`, `DEFAULT`, `MANAGED_DEVICE_HOSTNAME`
- `versioned_template_id` (String) Versioned template ID to be provisioned

Optional:

- `host_name` (String) Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
- `id` (String) ID of device is required if `type` is MANAGED_DEVICE_UUID
- `params` (Map of String) Template params/values to be provisioned
- `resource_params` (Map of String) Resource params to be provisioned


<a id="nestedatt--member_template_deployment_info"></a>
### Nested Schema for `member_template_deployment_info`

Required:

- `target_info` (Attributes List) Target info to deploy template (see [below for nested schema](#nestedatt--member_template_deployment_info--target_info))
- `template_id` (String) ID of template to be provisioned

Optional:

- `force_push_template` (Boolean) Force Push Template
- `is_composite` (Boolean) Composite template flag
- `main_template_id` (String) Composite Template ID

<a id="nestedatt--member_template_deployment_info--target_info"></a>
### Nested Schema for `member_template_deployment_info.target_info`

Required:

- `type` (String) Target type of device
  - Choices: `MANAGED_DEVICE_IP`, `MANAGED_DEVICE_UUID`, `PRE_PROVISIONED_SERIAL`, `PRE_PROVISIONED_MAC`, `DEFAULT`, `MANAGED_DEVICE_HOSTNAME`
- `versioned_template_id` (String) Versioned template ID to be provisioned

Optional:

- `host_name` (String) Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
- `id` (String) ID of device is required if targetType is MANAGED_DEVICE_UUID
- `params` (Map of String) Template params/values to be provisioned
- `resource_params` (Map of String) Resource params to be provisioned