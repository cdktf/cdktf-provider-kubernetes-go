// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicyv1


type NetworkPolicyV1SpecIngress struct {
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/network_policy_v1#from NetworkPolicyV1#from}
	From interface{} `field:"optional" json:"from" yaml:"from"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/network_policy_v1#ports NetworkPolicyV1#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

