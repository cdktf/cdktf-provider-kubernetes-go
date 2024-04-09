// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podsecuritypolicy


type PodSecurityPolicySpecRunAsGroup struct {
	// rule is the strategy that will dictate the allowable RunAsGroup values that may be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod_security_policy#rule PodSecurityPolicy#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/pod_security_policy#range PodSecurityPolicy#range}
	Range interface{} `field:"optional" json:"range" yaml:"range"`
}

