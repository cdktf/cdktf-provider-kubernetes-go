// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type JobSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#container_name Job#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#resource Job#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/job#divisor Job#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}
