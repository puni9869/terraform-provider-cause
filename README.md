# Terraform Provider Cause

Run the following command to build the provider

```shell
go build -o terraform-provider-cause
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
export TF_LOG=VERBOSE
cd examples && terraform init && terraform apply
```

[More](https://developer.hashicorp.com/terraform/plugin)