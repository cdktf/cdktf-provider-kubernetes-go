package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecInitContainerStartupProbeTcpSocket struct {
	// Number or name of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/replication_controller_v1#port ReplicationControllerV1#port}
	Port *string `field:"required" json:"port" yaml:"port"`
}

