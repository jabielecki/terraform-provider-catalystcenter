resource "catalystcenter_image_activation" "example" {
  # for_each = {
  #   for k, v in local.devices : k => v
  #   if v.platform_id == "C9407R"
  # }

  # device_uuid     = each.key
  device_uuid = "5b879fa4-c4d9-4d11-9adf-7af40e83f7a0"

  # image_uuid_list = [catalystcenter_image.example.id]
  image_uuid_list = ["8062b96b-c647-4f0f-a066-420c405bcb3c"] #  8062b96b-c647-4f0f-a066-420c405bcb3c the original throttled experimental software

  activate_lower_image_version = true
  distribute_if_needed         = true
}
