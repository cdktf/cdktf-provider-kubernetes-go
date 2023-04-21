package networkpolicy


type NetworkPolicySpecIngressFrom struct {
	// ip_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/network_policy#ip_block NetworkPolicy#ip_block}
	IpBlock *NetworkPolicySpecIngressFromIpBlock `field:"optional" json:"ipBlock" yaml:"ipBlock"`
	// namespace_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/network_policy#namespace_selector NetworkPolicy#namespace_selector}
	NamespaceSelector *NetworkPolicySpecIngressFromNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
	// pod_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/network_policy#pod_selector NetworkPolicy#pod_selector}
	PodSelector *NetworkPolicySpecIngressFromPodSelector `field:"optional" json:"podSelector" yaml:"podSelector"`
}

