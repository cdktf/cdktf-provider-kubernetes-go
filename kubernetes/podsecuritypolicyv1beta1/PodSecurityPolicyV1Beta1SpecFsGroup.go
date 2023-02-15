package podsecuritypolicyv1beta1


type PodSecurityPolicyV1Beta1SpecFsGroup struct {
	// rule is the strategy that will dictate what FSGroup is used in the SecurityContext.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy_v1beta1#rule PodSecurityPolicyV1Beta1#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy_v1beta1#range PodSecurityPolicyV1Beta1#range}
	Range interface{} `field:"optional" json:"range" yaml:"range"`
}

