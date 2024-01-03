resource "catalystcenter_discovery" "example" {
  cdp_level                 = 10
  name                      = "disco42"
  preferred_ip_method       = "UseLoopBack"
  discovery_type            = "RANGE"
  ip_address_list           = "192.168.0.1-192.168.0.99"
  ip_filter_list            = ["192.168.8.8-192.168.8.8"]
  global_credential_id_list = [""]
  protocol_order            = "ssh"
  netconf_port              = ""
}
