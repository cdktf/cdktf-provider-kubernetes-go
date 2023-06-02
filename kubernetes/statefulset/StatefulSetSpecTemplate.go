package statefulset


type StatefulSetSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#metadata StatefulSet#metadata}
	Metadata *StatefulSetSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.0/docs/resources/stateful_set#spec StatefulSet#spec}
	Spec *StatefulSetSpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

