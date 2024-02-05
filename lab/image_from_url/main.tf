locals {
  source_url = ""
}

resource "catalystcenter_image_from_url" "example" {
  // source_url  = "https://download-ssc.cisco.com/files/swc/sec/4_SDSP_16588806_1679981337552/1/cat9k_iosxe.17.11.01.SPA.bin?ip=10.36.7.13&dtrTag=4_SDSP_16588806_1679981337552_b87c30f357d0e049ee4306d2a34aaf34&userid=jabielec&ak-org=2w_ue&SEC=Y&tenant-id=SDSP&client=NA&__gda__=1707208097_3fe9d3fdce4919612fbbb93bcf54caf697b7d87f9b1a74a556616ab4d32636d9&__gdb__=exp=1707204978~hmac=9ab8bb7d3ded4c9b46204a78787fc0d102ffa89185f85daa1c50d6213510aa18&xac=4-PhO248%2Fmf26H3jjM%2BgBZdStG9Vs61ZEwxgfVCIEuShKDJIKIuQhuaETfCsUWmlBD"
  source_url  = "https://download-ssc.cisco.com/files/swc/sec/4_SDSP_16588806_1679981337552/1/cat9k_iosxe.17.11.01.SPA.bin"
  name        = "cat9k_iosxe.17.11.01.SPA.bin"
  vendor      = "CISCO" // FIXME
  family      = "CISCO" // FIXME
  third_party = false
}
