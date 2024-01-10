resource "catalystcenter_image_from_file" "example" {
  source_path = "../software.bin"
  name        = "basename(\"../software.bin\")"
}
