// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecOs struct {
	// Name is the name of the operating system. The currently supported values are linux and windows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/pod#name Pod#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

