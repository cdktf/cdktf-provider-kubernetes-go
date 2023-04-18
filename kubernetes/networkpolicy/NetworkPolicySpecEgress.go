package networkpolicy


type NetworkPolicySpecEgress struct {
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/network_policy#ports NetworkPolicy#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/network_policy#to NetworkPolicy#to}
	To interface{} `field:"optional" json:"to" yaml:"to"`
}

