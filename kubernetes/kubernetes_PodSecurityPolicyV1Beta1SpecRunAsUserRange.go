// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSecurityPolicyV1Beta1SpecRunAsUserRange struct {
	// max is the end of the range, inclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy_v1beta1#max PodSecurityPolicyV1Beta1#max}
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// min is the start of the range, inclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy_v1beta1#min PodSecurityPolicyV1Beta1#min}
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

