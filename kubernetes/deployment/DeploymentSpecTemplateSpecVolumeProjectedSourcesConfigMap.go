// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecVolumeProjectedSourcesConfigMap struct {
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#items Deployment#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#name Deployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional: Specify whether the ConfigMap or it's keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/deployment#optional Deployment#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

