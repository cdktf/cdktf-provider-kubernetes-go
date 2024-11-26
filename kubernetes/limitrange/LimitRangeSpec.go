// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package limitrange


type LimitRangeSpec struct {
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/limit_range#limit LimitRange#limit}
	Limit interface{} `field:"optional" json:"limit" yaml:"limit"`
}

