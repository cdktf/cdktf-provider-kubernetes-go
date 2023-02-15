package deployment


type DeploymentSpecTemplateSpecInitContainerPort struct {
	// Number of port to expose on the pod's IP address.
	//
	// This must be a valid port number, 0 < x < 65536.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#container_port Deployment#container_port}
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// What host IP to bind the external port to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#host_ip Deployment#host_ip}
	HostIp *string `field:"optional" json:"hostIp" yaml:"hostIp"`
	// Number of port to expose on the host.
	//
	// If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#host_port Deployment#host_port}
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// If specified, this must be an IANA_SVC_NAME and unique within the pod.
	//
	// Each named port in a pod must have a unique name. Name for the port that can be referred to by services
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#name Deployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Protocol for port. Must be UDP or TCP. Defaults to "TCP".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/deployment#protocol Deployment#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

