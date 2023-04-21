package jobv1


type JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItems struct {
	// Path is the relative path name of the file to be created.
	//
	// Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#path JobV1#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#field_ref JobV1#field_ref}
	FieldRef *JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// Mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#mode JobV1#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#resource_field_ref JobV1#resource_field_ref}
	ResourceFieldRef *JobV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

