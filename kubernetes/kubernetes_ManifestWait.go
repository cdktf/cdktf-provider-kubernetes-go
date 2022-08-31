// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ManifestWait struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#condition Manifest#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// A map of paths to fields to wait for a specific field value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#fields Manifest#fields}
	Fields *map[string]*string `field:"optional" json:"fields" yaml:"fields"`
	// Wait for rollout to complete on resources that support `kubectl rollout status`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/manifest#rollout Manifest#rollout}
	Rollout interface{} `field:"optional" json:"rollout" yaml:"rollout"`
}

