// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod#post_start Pod#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod#pre_stop Pod#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

