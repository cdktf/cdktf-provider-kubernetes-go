package networkpolicyv1


type NetworkPolicyV1SpecEgressPorts struct {
	// The port on the given protocol.
	//
	// This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/network_policy_v1#port NetworkPolicyV1#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/network_policy_v1#protocol NetworkPolicyV1#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

