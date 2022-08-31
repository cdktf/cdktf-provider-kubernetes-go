// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type DaemonSetV1SpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#metadata DaemonSetV1#metadata}
	Metadata *DaemonSetV1SpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#spec DaemonSetV1#spec}
	Spec *DaemonSetV1SpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

