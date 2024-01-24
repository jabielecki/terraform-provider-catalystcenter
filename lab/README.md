# Lab Deployment

## Install Terraform

```sh
curl -L https://releases.hashicorp.com/terraform/1.6.6/terraform_1.6.6_linux_amd64.zip | gunzip | sudo tee /usr/local/bin/terraform >/dev/null && sudo chmod 0755 /usr/local/bin/terraform
```

## Double Check Code Generation

Some code changes require a generator to be run. It doesn't hurt to do it again, just to be sure:

```sh
go generate .
```

## Run Terraform

Replace the env here to point to your actual Cisco Catalyst Center:

```sh
export CC_USERNAME=admin CC_PASSWORD='secret' CC_URL=https://10.1.1.1
export TF_LOG=debug TF_LOG_PATH=log.txt
```

Here, enter the scenario (one of `lab/*` directories):

```sh
BI="$PWD/bin"

cd lab/network_devices

rm -f log.txt

CGO_ENABLED=0 go build -o "$BI"/registry.terraform.io/CiscoDevNet/catalystcenter/0.0.9999/linux_amd64/terraform-provider-catalystcenter "$BI/../"

terraform providers lock -fs-mirror="$BI"

terraform init -plugin-dir="$BI"

terraform apply -auto-approve

terraform destroy -auto-approve

egrep 'HTTP Re|panic|swim' log.txt

cd -
```

All these commands are suitable to run as a single loooong command line, or as a shell alias.

## Explanation

The default CGo compilation is disabled, which results in a binary more compatible with comparatively more Linux distros, old and new, in case you move the newly built binary to another Linux system.

The binary location in its directory hierarchy is critical so that Terraform could pick it up from disk (otherwise Terraform would attempt to download from the official Registry, and fail). In particular, the `.../CiscoDevNet/catalystcenter/...` matches the version in the supplied `lab/*/provider.tf`:

```hcl2
terraform {
  required_providers {
    catalystcenter = {
      source  = "CiscoDevNet/catalystcenter"
    }
  }
}
```
