resource "catalystcenter_ip_pool" "test" {
  name      = "MyPool1"
  ip_subnet = "172.32.0.0/16"
}
