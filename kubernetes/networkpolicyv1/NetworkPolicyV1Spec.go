package networkpolicyv1


type NetworkPolicyV1Spec struct {
	// pod_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/network_policy_v1#pod_selector NetworkPolicyV1#pod_selector}
	PodSelector *NetworkPolicyV1SpecPodSelector `field:"required" json:"podSelector" yaml:"podSelector"`
	// List of rule types that the NetworkPolicy relates to.
	//
	// Valid options are ["Ingress"], ["Egress"], or ["Ingress", "Egress"]. If this field is not specified, it will default based on the existence of Ingress or Egress rules; policies that contain an Egress section are assumed to affect Egress, and all policies (whether or not they contain an Ingress section) are assumed to affect Ingress. If you want to write an egress-only policy, you must explicitly specify policyTypes [ "Egress" ]. Likewise, if you want to write a policy that specifies that no egress is allowed, you must specify a policyTypes value that include "Egress" (since such a policy would not include an Egress section and would otherwise default to just [ "Ingress" ]). This field is beta-level in 1.8
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/network_policy_v1#policy_types NetworkPolicyV1#policy_types}
	PolicyTypes *[]*string `field:"required" json:"policyTypes" yaml:"policyTypes"`
	// egress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/network_policy_v1#egress NetworkPolicyV1#egress}
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// ingress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/network_policy_v1#ingress NetworkPolicyV1#ingress}
	Ingress interface{} `field:"optional" json:"ingress" yaml:"ingress"`
}

