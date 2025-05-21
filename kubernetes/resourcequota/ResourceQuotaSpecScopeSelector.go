// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcequota


type ResourceQuotaSpecScopeSelector struct {
	// match_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.0/docs/resources/resource_quota#match_expression ResourceQuota#match_expression}
	MatchExpression interface{} `field:"optional" json:"matchExpression" yaml:"matchExpression"`
}

