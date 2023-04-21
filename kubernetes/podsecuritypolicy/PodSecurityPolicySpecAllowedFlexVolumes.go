package podsecuritypolicy


type PodSecurityPolicySpecAllowedFlexVolumes struct {
	// driver is the name of the Flexvolume driver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod_security_policy#driver PodSecurityPolicy#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
}

