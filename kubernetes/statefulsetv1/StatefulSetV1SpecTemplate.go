package statefulsetv1


type StatefulSetV1SpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#metadata StatefulSetV1#metadata}
	Metadata *StatefulSetV1SpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#spec StatefulSetV1#spec}
	Spec *StatefulSetV1SpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

