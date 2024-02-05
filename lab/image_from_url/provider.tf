terraform {
  required_providers {
    catalystcenter = {
      source  = "CiscoDevNet/catalystcenter"
      version = "0.0.9999"
    }
  }
}

provider "catalystcenter" {
  // password = "plaintext" // provided instead with $CC_PASSWORD
}
