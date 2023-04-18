package daemonset


type DaemonsetSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemonset#metadata Daemonset#metadata}
	Metadata *DaemonsetSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/daemonset#spec Daemonset#spec}
	Spec *DaemonsetSpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

