package replicationcontroller


type ReplicationControllerSpecTemplateSpecSecurityContextSysctl struct {
	// Name of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#name ReplicationController#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#value ReplicationController#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

