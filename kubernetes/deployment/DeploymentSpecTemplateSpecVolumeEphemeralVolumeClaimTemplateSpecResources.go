// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecResources struct {
	// Map describing the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/deployment#limits Deployment#limits}
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
	// Map describing the minimum amount of compute resources required.
	//
	// If this is omitted for a container, it defaults to `limits` if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/deployment#requests Deployment#requests}
	Requests *map[string]*string `field:"optional" json:"requests" yaml:"requests"`
}

