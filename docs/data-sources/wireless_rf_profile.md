---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_rf_profile Data Source - terraform-provider-catalystcenter"
subcategory: "Wireless"
description: |-
  This data source can read the Wireless RF Profile.
---

# catalystcenter_wireless_rf_profile (Data Source)

This data source can read the Wireless RF Profile.

## Example Usage

```terraform
data "catalystcenter_wireless_rf_profile" "example" {
  id = "RF_Profile_1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `channel_width` (String) Channel Width
- `default_rf_profile` (Boolean) is Default Rf Profile
- `enable_brown_field` (Boolean) Enable Brown Field
- `enable_custom` (Boolean) Enable Custom
- `enable_radio_type_a` (Boolean) Enable Radio Type A
- `enable_radio_type_b` (Boolean) Enable Radio Type B
- `enable_radio_type_c` (Boolean) Enable Radio Type C (6GHz)
- `name` (String) RF Profile Name
- `radio_type_a_data_rates` (String) Radio TypeA Properties - Data Rates
- `radio_type_a_mandatory_data_rates` (String) Radio TypeA Properties - Mandatory Data Rates
- `radio_type_a_max_power_level` (Number) Radio TypeA Properties - Max Power Level
- `radio_type_a_min_power_level` (Number) Radio TypeA Properties - Min Power Level
- `radio_type_a_parent_profile` (String) Radio TypeA Properties - Parent Profile
- `radio_type_a_power_treshold_v1` (Number) Radio TypeA Properties - Power Threshold V1
- `radio_type_a_radio_channels` (String) Radio TypeA Properties - Radio Channels
- `radio_type_a_rx_sop_threshold` (String) Radio TypeA Properties - Rx Sop Threshold
- `radio_type_b_data_rates` (String) Radio TypeB Properties - Data Rates
- `radio_type_b_mandatory_data_rates` (String) Radio TypeB Properties - Mandatory Data Rates
- `radio_type_b_max_power_level` (Number) Radio TypeB Properties - Max Power Level
- `radio_type_b_min_power_level` (Number) Radio TypeB Properties - Min Power Level
- `radio_type_b_parent_profile` (String) Radio TypeB Properties - Parent Profile
- `radio_type_b_power_treshold_v1` (Number) Radio TypeB Properties - Power Threshold V1
- `radio_type_b_radio_channels` (String) Radio TypeB Properties - Radio Channels
- `radio_type_b_rx_sop_threshold` (String) Radio TypeB Properties - Rx Sop Threshold
- `radio_type_c_data_rates` (String) Radio TypeC Properties - Data Rates
- `radio_type_c_mandatory_data_rates` (String) Radio TypeC Properties - Mandatory Data Rates
- `radio_type_c_max_power_level` (Number) Radio TypeC Properties - Max Power Level
- `radio_type_c_min_power_level` (Number) Radio TypeC Properties - Min Power Level
- `radio_type_c_parent_profile` (String) Radio TypeC Properties - Parent Profile
- `radio_type_c_power_treshold_v1` (Number) Radio TypeC Properties - Power Threshold V1
- `radio_type_c_radio_channels` (String) Radio TypeC Properties - Radio Channels
- `radio_type_c_rx_sop_threshold` (String) Radio TypeC Properties - Rx Sop Threshold