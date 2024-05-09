// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicy


type NetworkPolicySpec struct {
	// pod_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/network_policy#pod_selector NetworkPolicy#pod_selector}
	PodSelector *NetworkPolicySpecPodSelector `field:"required" json:"podSelector" yaml:"podSelector"`
	// policyTypes is a list of rule types that the NetworkPolicy relates to.
	//
	// Valid options are ["Ingress"], ["Egress"], or ["Ingress", "Egress"]. If this field is not specified, it will default based on the existence of ingress or egress rules; policies that contain an egress section are assumed to affect egress, and all policies (whether or not they contain an ingress section) are assumed to affect ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/network_policy#policy_types NetworkPolicy#policy_types}
	PolicyTypes *[]*string `field:"required" json:"policyTypes" yaml:"policyTypes"`
	// egress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/network_policy#egress NetworkPolicy#egress}
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// ingress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/network_policy#ingress NetworkPolicy#ingress}
	Ingress interface{} `field:"optional" json:"ingress" yaml:"ingress"`
}

