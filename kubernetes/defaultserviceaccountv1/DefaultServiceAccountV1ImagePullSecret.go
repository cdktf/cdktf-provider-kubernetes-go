// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package defaultserviceaccountv1


type DefaultServiceAccountV1ImagePullSecret struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/default_service_account_v1#name DefaultServiceAccountV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

