package statefulset


type StatefulSetSpecVolumeClaimTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/stateful_set#metadata StatefulSet#metadata}
	Metadata *StatefulSetSpecVolumeClaimTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/stateful_set#spec StatefulSet#spec}
	Spec *StatefulSetSpecVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

