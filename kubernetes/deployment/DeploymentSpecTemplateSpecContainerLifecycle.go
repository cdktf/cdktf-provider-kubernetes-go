// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/deployment#post_start Deployment#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/deployment#pre_stop Deployment#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

