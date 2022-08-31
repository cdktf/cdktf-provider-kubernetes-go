// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ManifestTimeouts struct {
	// Timeout for the create operation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#create Manifest#create}
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Timeout for the delete operation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#delete Manifest#delete}
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Timeout for the update operation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#update Manifest#update}
	Update *string `field:"optional" json:"update" yaml:"update"`
}

