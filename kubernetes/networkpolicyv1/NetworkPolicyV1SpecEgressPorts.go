// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkpolicyv1


type NetworkPolicyV1SpecEgressPorts struct {
	// endPort indicates that the range of ports from port to endPort if set, inclusive, should be allowed by the policy.
	//
	// This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/network_policy_v1#end_port NetworkPolicyV1#end_port}
	EndPort *float64 `field:"optional" json:"endPort" yaml:"endPort"`
	// port represents the port on the given protocol.
	//
	// This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/network_policy_v1#port NetworkPolicyV1#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match.
	//
	// If not specified, this field defaults to TCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/network_policy_v1#protocol NetworkPolicyV1#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

