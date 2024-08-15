# CDKTF Go bindings for hashicorp/kubernetes provider version 2.32.0

This repo builds and publishes the [Terraform kubernetes provider](https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs) bindings for [CDK for Terraform](https://cdk.tf).

## Go Package

The go package is generated into the [`github.com/cdktf/cdktf-provider-kubernetes-go`](https://github.com/cdktf/cdktf-provider-kubernetes-go) package.

`go get github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/<version>`

Where `<version>` is the version of the prebuilt provider you would like to use e.g. `v11`. The full module name can be found
within the [go.mod](https://github.com/cdktf/cdktf-provider-kubernetes-go/blob/main/kubernetes/go.mod#L1) file.

## Docs

Find auto-generated docs for this provider [here](https://github.com/cdktf/cdktf-provider-kubernetes/blob/main/docs/API.go.md).


## Versioning

This project is explicitly not tracking the Terraform kubernetes provider version 1:1. In fact, it always tracks `latest` of `~> 2.0` with every release. If there are scenarios where you explicitly have to pin your provider version, you can do so by [generating the provider constructs manually](https://cdk.tf/imports).

These are the upstream dependencies:

* [CDK for Terraform](https://cdk.tf)
* [Terraform kubernetes provider](https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0)
* [Terraform Engine](https://terraform.io)

If there are breaking changes (backward incompatible) in any of the above, the major version of this project will be bumped.

## Features / Issues / Bugs

Please report bugs and issues to the [CDK for Terraform](https://cdk.tf) project:

* [Create bug report](https://cdk.tf/bug)
* [Create feature request](https://cdk.tf/feature)

## Contributing

### Projen

This is mostly based on [Projen](https://github.com/projen/projen), which takes care of generating the entire repository.

### cdktf-provider-project based on Projen

There's a custom [project builder](https://github.com/cdktf/cdktf-provider-project) which encapsulate the common settings for all `cdktf` prebuilt providers.


### Repository Management

The repository is managed by [CDKTF Repository Manager](https://github.com/cdktf/cdktf-repository-manager/).
