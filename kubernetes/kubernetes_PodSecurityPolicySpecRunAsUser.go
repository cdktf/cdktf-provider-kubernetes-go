// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSecurityPolicySpecRunAsUser struct {
	// rule is the strategy that will dictate the allowable RunAsUser values that may be set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy#rule PodSecurityPolicy#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// range block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy#range PodSecurityPolicy#range}
	Range interface{} `field:"optional" json:"range" yaml:"range"`
}

