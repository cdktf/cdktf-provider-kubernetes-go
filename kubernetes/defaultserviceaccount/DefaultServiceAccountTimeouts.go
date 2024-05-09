// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaultserviceaccount


type DefaultServiceAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/default_service_account#create DefaultServiceAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

