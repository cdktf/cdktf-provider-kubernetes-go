// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type NamespaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/namespace#delete Namespace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

