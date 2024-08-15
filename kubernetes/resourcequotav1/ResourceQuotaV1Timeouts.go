// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcequotav1


type ResourceQuotaV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/resource_quota_v1#create ResourceQuotaV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/resource_quota_v1#update ResourceQuotaV1#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

