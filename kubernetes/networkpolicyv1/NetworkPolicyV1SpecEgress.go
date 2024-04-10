// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicyv1


type NetworkPolicyV1SpecEgress struct {
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/network_policy_v1#ports NetworkPolicyV1#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.1/docs/resources/network_policy_v1#to NetworkPolicyV1#to}
	To interface{} `field:"optional" json:"to" yaml:"to"`
}

