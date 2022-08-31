// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type NetworkPolicySpecEgress struct {
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/network_policy#ports NetworkPolicy#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/network_policy#to NetworkPolicy#to}
	To interface{} `field:"optional" json:"to" yaml:"to"`
}

