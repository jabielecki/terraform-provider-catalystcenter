locals {
  url = "http://ftpmirror.uk/pub/Software/Cisco/3850/cat3k_caa-universalk9.16.12.07.SPA.bin"
}
resource "catalystcenter_image_from_url" "example" {
  source_url = local.url
  name       = basename(local.url)
}

output "catalystcenter_image_from_url" {
  value = catalystcenter_image_from_url.example
}
