locals {
  source_paths = sort(tolist(fileset(path.module, "*.bin")))
  source_path  = try(local.source_paths[0], "")
}

# resource "catalystcenter_image" "example" {

#   source_path = local.source_path
#   name        = basename(local.source_path)

#   lifecycle {
#     precondition {
#       condition     = length(fileset(path.module, "*.bin")) > 0
#       error_message = "Software image file not found! Please put it here: ${abspath(path.module)}/*.bin"
#     }
#   }

# }

# output "catalystcenter_image" {
#   value = catalystcenter_image.example
# }

resource "catalystcenter_image" "test" {
  third_party_application_type = "UNKNOWN"
  family                       = "TF"
  source_path                  = "/tmp/TF.1.2.3-comp_matrix.zip"
  name                         = basename("/tmp/TF.1.2.3-comp_matrix.zip")
  third_party_vendor           = "CISCO"
  is_third_party               = true
}
