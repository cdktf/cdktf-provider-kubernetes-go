// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSpecContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#add Pod#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#drop Pod#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

