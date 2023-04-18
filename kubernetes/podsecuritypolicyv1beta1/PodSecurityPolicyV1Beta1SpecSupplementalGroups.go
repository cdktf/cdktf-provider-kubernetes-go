package podsecuritypolicyv1beta1


type PodSecurityPolicyV1Beta1SpecSupplementalGroups struct {
	// rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/pod_security_policy_v1beta1#rule PodSecurityPolicyV1Beta1#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/pod_security_policy_v1beta1#range PodSecurityPolicyV1Beta1#range}
	Range interface{} `field:"optional" json:"range" yaml:"range"`
}

