// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podsecuritypolicy


type PodSecurityPolicySpecRunAsUserRange struct {
	// max is the end of the range, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod_security_policy#max PodSecurityPolicy#max}
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// min is the start of the range, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod_security_policy#min PodSecurityPolicy#min}
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

