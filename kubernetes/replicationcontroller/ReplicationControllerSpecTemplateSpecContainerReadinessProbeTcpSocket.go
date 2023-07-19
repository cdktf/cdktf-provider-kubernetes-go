package replicationcontroller


type ReplicationControllerSpecTemplateSpecContainerReadinessProbeTcpSocket struct {
	// Number or name of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.22.0/docs/resources/replication_controller#port ReplicationController#port}
	Port *string `field:"required" json:"port" yaml:"port"`
}

