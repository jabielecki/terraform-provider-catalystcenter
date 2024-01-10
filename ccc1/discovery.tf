# resource "catalystcenter_discovery" "example" {
#   discovery_type            = "RANGE"
#   global_credential_id_list = ["a535265c-8c96-498b-ba9f-2a94c9e1f806"] // uuid manually copied from a CLI credential
#   ip_address_list           = "204.101.16.1-204.101.16.9"
#   name                      = "disco42"
#   netconf_port              = "830"
#   preferred_ip_method       = "UseLoopBack"
#   protocol_order            = "ssh"
#   retry                     = 3
#   timeout                   = 5
# }

# # In fake-POST scenario, terraform refresh generally does not update these attributes, but terraform delete does (just before deleting them).
# # output "catalystcenter_discovery" {
# #   value = catalystcenter_discovery.example
# # }

# data "catalystcenter_discovery" "example" {
#   id = catalystcenter_discovery.example.id
# }

# output "data_catalystcenter_discovery" {
#   value = data.catalystcenter_discovery.example
# }
