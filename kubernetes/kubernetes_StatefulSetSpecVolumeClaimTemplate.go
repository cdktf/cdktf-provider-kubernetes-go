// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type StatefulSetSpecVolumeClaimTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#metadata StatefulSet#metadata}
	Metadata *StatefulSetSpecVolumeClaimTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#spec StatefulSet#spec}
	Spec *StatefulSetSpecVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

