// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secret


type SecretTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/secret#create Secret#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

