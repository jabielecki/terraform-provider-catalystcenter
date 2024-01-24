locals {
  source_paths = sort(tolist(fileset(path.module, "*.bin")))
  source_path  = try(local.source_paths[0], "")
}

resource "catalystcenter_image" "example" {

  source_path = local.source_path
  name        = basename(local.source_path)

  lifecycle {
    precondition {
      condition     = length(fileset(path.module, "*.bin")) > 0
      error_message = "Software image file not found! Please put it here: ${abspath(path.module)}/*.bin"
    }
  }

}

output "catalystcenter_image" {
  value = catalystcenter_image.example
}
