// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceaccountv1


type ServiceAccountV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.27.0/docs/resources/service_account_v1#create ServiceAccountV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

