package statefulset


type StatefulSetSpecTemplateSpecContainerLifecyclePostStartTcpSocket struct {
	// Number or name of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/stateful_set#port StatefulSet#port}
	Port *string `field:"required" json:"port" yaml:"port"`
}

