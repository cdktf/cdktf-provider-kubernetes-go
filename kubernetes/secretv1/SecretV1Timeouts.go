// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretv1


type SecretV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/secret_v1#create SecretV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

