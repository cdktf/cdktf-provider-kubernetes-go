// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ManifestWaitFor struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#fields Manifest#fields}.
	Fields *map[string]*string `field:"optional" json:"fields" yaml:"fields"`
}

