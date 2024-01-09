// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcequotav1


type ResourceQuotaV1SpecScopeSelectorMatchExpression struct {
	// Represents a scope's relationship to a set of values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/resource_quota_v1#operator ResourceQuotaV1#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The name of the scope that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/resource_quota_v1#scope_name ResourceQuotaV1#scope_name}
	ScopeName *string `field:"required" json:"scopeName" yaml:"scopeName"`
	// A list of scope selector requirements by scope of the resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/resource_quota_v1#values ResourceQuotaV1#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

