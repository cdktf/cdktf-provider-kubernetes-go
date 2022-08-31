// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSecurityPolicySpecAllowedFlexVolumes struct {
	// driver is the name of the Flexvolume driver.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod_security_policy#driver PodSecurityPolicy#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
}

