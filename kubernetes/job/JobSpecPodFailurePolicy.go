// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecPodFailurePolicy struct {
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/job#rule Job#rule}
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
}

