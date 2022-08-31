// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type StatefulSetV1SpecVolumeClaimTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#metadata StatefulSetV1#metadata}
	Metadata *StatefulSetV1SpecVolumeClaimTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set_v1#spec StatefulSetV1#spec}
	Spec *StatefulSetV1SpecVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

