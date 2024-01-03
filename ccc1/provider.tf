terraform {
  required_providers {
    catalystcenter = {
      source = "codilime.com/CiscoDevNet/catalystcenter"
      version = "0.1.1"
    }
  }
}

provider "catalystcenter" {
  // password = "plaintext" // provided instead with $CC_PASSWORD
}
