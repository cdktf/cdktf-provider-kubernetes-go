package replicationcontrollerv1


type ReplicationControllerV1SpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#metadata ReplicationControllerV1#metadata}
	Metadata *ReplicationControllerV1SpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller_v1#spec ReplicationControllerV1#spec}
	Spec *ReplicationControllerV1SpecTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

