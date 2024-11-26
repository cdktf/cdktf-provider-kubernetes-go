// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package limitrange


type LimitRangeSpecLimit struct {
	// Default resource requirement limit value by resource name if resource limit is omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/limit_range#default LimitRange#default}
	Default *map[string]*string `field:"optional" json:"default" yaml:"default"`
	// The default resource requirement request value by resource name if resource request is omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/limit_range#default_request LimitRange#default_request}
	DefaultRequest *map[string]*string `field:"optional" json:"defaultRequest" yaml:"defaultRequest"`
	// Max usage constraints on this kind by resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/limit_range#max LimitRange#max}
	Max *map[string]*string `field:"optional" json:"max" yaml:"max"`
	// The named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value;
	//
	// this represents the max burst for the named resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/limit_range#max_limit_request_ratio LimitRange#max_limit_request_ratio}
	MaxLimitRequestRatio *map[string]*string `field:"optional" json:"maxLimitRequestRatio" yaml:"maxLimitRequestRatio"`
	// Min usage constraints on this kind by resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/limit_range#min LimitRange#min}
	Min *map[string]*string `field:"optional" json:"min" yaml:"min"`
	// Type of resource that this limit applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/limit_range#type LimitRange#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

