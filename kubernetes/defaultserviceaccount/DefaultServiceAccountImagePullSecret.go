// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaultserviceaccount


type DefaultServiceAccountImagePullSecret struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/default_service_account#name DefaultServiceAccount#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

