package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecDnsConfigOption struct {
	// Name of the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/replication_controller_v1#name ReplicationControllerV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of the option. Optional: Defaults to empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.19.0/docs/resources/replication_controller_v1#value ReplicationControllerV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}
