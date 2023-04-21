package podsecuritypolicyv1beta1


type PodSecurityPolicyV1Beta1SpecHostPorts struct {
	// max is the end of the range, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_security_policy_v1beta1#max PodSecurityPolicyV1Beta1#max}
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// min is the start of the range, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_security_policy_v1beta1#min PodSecurityPolicyV1Beta1#min}
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

