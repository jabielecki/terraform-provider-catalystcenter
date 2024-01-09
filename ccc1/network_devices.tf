data "catalystcenter_network_devices" "example" {
  id = "remove this attribute" // FIXME
}

locals {
  devices = {
    for v in data.catalystcenter_network_devices.example.response : v.id => {
      management_ip_address = v.management_ip_address
      platform_id           = v.platform_id
      role                  = v.role
      software_type         = v.software_type
      hostname              = v.hostname
    }
  }
}

output "catalystcenter_network_devices" {
  value = local.devices
}
