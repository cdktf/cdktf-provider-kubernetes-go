// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod_v1#name PodV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod_v1#value PodV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

