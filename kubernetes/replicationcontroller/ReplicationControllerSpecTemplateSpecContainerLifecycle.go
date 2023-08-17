package replicationcontroller


type ReplicationControllerSpecTemplateSpecContainerLifecycle struct {
	// post_start block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#post_start ReplicationController#post_start}
	PostStart interface{} `field:"optional" json:"postStart" yaml:"postStart"`
	// pre_stop block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/replication_controller#pre_stop ReplicationController#pre_stop}
	PreStop interface{} `field:"optional" json:"preStop" yaml:"preStop"`
}

