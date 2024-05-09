// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podsecuritypolicyv1beta1


type PodSecurityPolicyV1Beta1SpecSeLinuxSeLinuxOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/pod_security_policy_v1beta1#level PodSecurityPolicyV1Beta1#level}.
	Level *string `field:"required" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/pod_security_policy_v1beta1#role PodSecurityPolicyV1Beta1#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/pod_security_policy_v1beta1#type PodSecurityPolicyV1Beta1#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/pod_security_policy_v1beta1#user PodSecurityPolicyV1Beta1#user}.
	User *string `field:"required" json:"user" yaml:"user"`
}

