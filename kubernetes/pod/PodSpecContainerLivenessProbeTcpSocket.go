package pod


type PodSpecContainerLivenessProbeTcpSocket struct {
	// Number or name of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/pod#port Pod#port}
	Port *string `field:"required" json:"port" yaml:"port"`
}

