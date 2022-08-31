// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodSpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#name Pod#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#value Pod#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

