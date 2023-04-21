package job


type JobSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#metadata Job#metadata}
	Metadata *JobSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#spec Job#spec}
	Spec *JobSpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

