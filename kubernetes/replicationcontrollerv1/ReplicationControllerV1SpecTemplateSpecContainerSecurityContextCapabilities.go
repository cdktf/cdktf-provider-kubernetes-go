package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/replication_controller_v1#add ReplicationControllerV1#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/replication_controller_v1#drop ReplicationControllerV1#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

