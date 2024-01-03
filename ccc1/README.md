# Test Deployment

## Install Terraform

```sh
curl -L https://releases.hashicorp.com/terraform/1.6.6/terraform_1.6.6_linux_amd64.zip | gunzip | sudo tee /usr/local/bin/terraform >/dev/null && sudo chmod 0755 /usr/local/bin/terraform
```

## Run Terraform

Unpack to an *empty* dir, without any pre-existing files from older versions.

Replace the env here to point to your actual Cisco Catalyst Center:

```sh
export CC_URL=https://10.1.1.1
export CC_USERNAME=admin
export CC_PASSWORD=secret

export TF_LOG=debug TF_LOG_PROVIDER=debug TF_LOG_PATH=log.txt

terraform init

terraform apply -auto-approve && terraform destroy -auto-approve
```

Please send the file `log.txt` back.
