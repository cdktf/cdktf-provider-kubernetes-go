package replicationcontroller


type ReplicationControllerSpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#metadata ReplicationController#metadata}
	Metadata *ReplicationControllerSpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#spec ReplicationController#spec}
	Spec *ReplicationControllerSpecTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

