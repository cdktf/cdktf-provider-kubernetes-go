package networkpolicyv1


type NetworkPolicyV1SpecIngress struct {
	// from block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/network_policy_v1#from NetworkPolicyV1#from}
	From interface{} `field:"optional" json:"from" yaml:"from"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/network_policy_v1#ports NetworkPolicyV1#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

