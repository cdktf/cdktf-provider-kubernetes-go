package podsecuritypolicy


type PodSecurityPolicySpecHostPorts struct {
	// max is the end of the range, inclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy#max PodSecurityPolicy#max}
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// min is the start of the range, inclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy#min PodSecurityPolicy#min}
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

