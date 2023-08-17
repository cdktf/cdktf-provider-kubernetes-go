package job


type JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate struct {
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/job#spec Job#spec}
	Spec *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/job#metadata Job#metadata}
	Metadata *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

