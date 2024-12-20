// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecContainerSecurityContextSeccompProfile struct {
	// Localhost Profile indicates a profile defined in a file on the node should be used.
	//
	// The profile must be preconfigured on the node to work.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_v1#localhost_profile PodV1#localhost_profile}
	LocalhostProfile *string `field:"optional" json:"localhostProfile" yaml:"localhostProfile"`
	// Type indicates which kind of seccomp profile will be applied. Valid options are: Localhost, RuntimeDefault, Unconfined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/pod_v1#type PodV1#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

