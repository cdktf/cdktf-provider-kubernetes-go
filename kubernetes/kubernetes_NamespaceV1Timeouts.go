// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type NamespaceV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/namespace_v1#delete NamespaceV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

