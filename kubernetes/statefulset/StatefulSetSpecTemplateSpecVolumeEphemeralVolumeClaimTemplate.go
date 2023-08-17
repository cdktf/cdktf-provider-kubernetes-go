package statefulset


type StatefulSetSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate struct {
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#spec StatefulSet#spec}
	Spec *StatefulSetSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#metadata StatefulSet#metadata}
	Metadata *StatefulSetSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

