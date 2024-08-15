// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type KubernetesProviderExec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs#api_version KubernetesProvider#api_version}.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs#command KubernetesProvider#command}.
	Command *string `field:"required" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs#args KubernetesProvider#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs#env KubernetesProvider#env}.
	Env *map[string]*string `field:"optional" json:"env" yaml:"env"`
}

