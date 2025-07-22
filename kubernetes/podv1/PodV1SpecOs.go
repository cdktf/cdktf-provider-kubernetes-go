// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecOs struct {
	// Name is the name of the operating system. The currently supported values are linux and windows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/pod_v1#name PodV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

