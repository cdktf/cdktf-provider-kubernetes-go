// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podsecuritypolicy


type PodSecurityPolicySpecSupplementalGroups struct {
	// rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/pod_security_policy#rule PodSecurityPolicy#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/pod_security_policy#range PodSecurityPolicy#range}
	Range interface{} `field:"optional" json:"range" yaml:"range"`
}

