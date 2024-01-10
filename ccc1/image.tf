locals {
  source_paths = sort(tolist(fileset(path.module, "*.bin")))
  source_path  = try(local.source_paths[0], "")
}

resource "catalystcenter_image_from_file" "example" {

  source_path = local.source_path
  name        = basename(local.source_path)

  lifecycle {
    precondition {
      condition     = length(fileset(path.module, "*.bin")) > 0
      error_message = "Software image file not found! Please put it here: ${abspath(path.module)}/*.bin"
    }
  }

}

output "catalystcenter_image_from_file" {
  value = catalystcenter_image_from_file.example
}

# resource "catalystcenter_image_distribution" "example" {
#   for_each = {
#     for k, v in local.devices : k => v
#     if v.role == "CORE"
#   }

#   device_uuid = each.key
#   image_uuid  = catalystcenter_image_from_file.example.id
# }
