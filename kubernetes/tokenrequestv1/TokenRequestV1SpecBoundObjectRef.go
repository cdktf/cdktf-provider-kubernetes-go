// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokenrequestv1


type TokenRequestV1SpecBoundObjectRef struct {
	// API version of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/token_request_v1#api_version TokenRequestV1#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the referent. Valid kinds are 'Pod' and 'Secret'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/token_request_v1#kind TokenRequestV1#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Name of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/token_request_v1#name TokenRequestV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// UID of the referent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/token_request_v1#uid TokenRequestV1#uid}
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

