resource "catalystcenter_image_distribution" "example" {
  # for_each = {
  #   for k, v in local.devices : k => v
  #   if v.hostname == "ISBTE-TB3-9400.cisco.com"
  # }

  # device_uuid = each.key
  device_uuid = "5b879fa4-c4d9-4d11-9adf-7af40e83f7a0"
  image_uuid  = "8062b96b-c647-4f0f-a066-420c405bcb3c"
}
