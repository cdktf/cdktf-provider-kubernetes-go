// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicev1


type ServiceV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.26.0/docs/resources/service_v1#create ServiceV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

