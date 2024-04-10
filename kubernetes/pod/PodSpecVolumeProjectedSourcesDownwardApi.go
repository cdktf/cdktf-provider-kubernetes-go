// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecVolumeProjectedSourcesDownwardApi struct {
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/pod#items Pod#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

