data "catalystcenter_device_detail" "example" {
  id = "5b879fa4-c4d9-4d11-9adf-7af40e83f7a0"
}

output "verify" {
  value = data.catalystcenter_device_detail.example
}
