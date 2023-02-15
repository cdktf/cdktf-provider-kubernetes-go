package networkpolicyv1


type NetworkPolicyV1SpecEgress struct {
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/network_policy_v1#ports NetworkPolicyV1#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/network_policy_v1#to NetworkPolicyV1#to}
	To interface{} `field:"optional" json:"to" yaml:"to"`
}

