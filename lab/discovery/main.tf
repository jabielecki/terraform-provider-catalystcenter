resource "catalystcenter_discovery" "example" {
  discovery_type            = "Range"
  global_credential_id_list = ["a535265c-8c96-498b-ba9f-2a94c9e1f806"] // uuid manually copied from a CLI credential
  ip_address_list           = "204.101.48.2-204.101.48.20"
  name                      = "disco42"
  netconf_port              = "830"
  protocol_order            = "SSH"
  retry                     = 3
  timeout_seconds           = 6
}

# output "catalystcenter_discovery" {
#   value = catalystcenter_discovery.example
# }
