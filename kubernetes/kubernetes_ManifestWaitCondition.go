// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ManifestWaitCondition struct {
	// The condition status.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#status Manifest#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The type of condition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#type Manifest#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

