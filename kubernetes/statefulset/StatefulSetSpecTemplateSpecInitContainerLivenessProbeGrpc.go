package statefulset


type StatefulSetSpecTemplateSpecInitContainerLivenessProbeGrpc struct {
	// Number of the port to access on the container. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#port StatefulSet#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/stateful_set#service StatefulSet#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

