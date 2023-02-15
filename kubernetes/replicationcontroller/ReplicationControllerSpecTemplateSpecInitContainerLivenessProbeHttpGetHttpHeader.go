package replicationcontroller


type ReplicationControllerSpecTemplateSpecInitContainerLivenessProbeHttpGetHttpHeader struct {
	// The header field name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#name ReplicationController#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/replication_controller#value ReplicationController#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

