// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceaccountv1


type ServiceAccountV1Secret struct {
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/service_account_v1#name ServiceAccountV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

