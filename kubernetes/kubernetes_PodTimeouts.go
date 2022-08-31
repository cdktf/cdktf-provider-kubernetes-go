// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type PodTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#create Pod#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/pod#delete Pod#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

