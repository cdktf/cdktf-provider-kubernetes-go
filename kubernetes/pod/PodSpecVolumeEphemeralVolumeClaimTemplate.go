package pod


type PodSpecVolumeEphemeralVolumeClaimTemplate struct {
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#spec Pod#spec}
	Spec *PodSpecVolumeEphemeralVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/pod#metadata Pod#metadata}
	Metadata *PodSpecVolumeEphemeralVolumeClaimTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

