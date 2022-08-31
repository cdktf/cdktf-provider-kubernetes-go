// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSecurityPolicyV1Beta1SpecAllowedFlexVolumes struct {
	// driver is the name of the Flexvolume driver.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy_v1beta1#driver PodSecurityPolicyV1Beta1#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
}

