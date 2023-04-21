package daemonset


type DaemonsetSpecTemplateSpecContainerLivenessProbeHttpGet struct {
	// Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#host Daemonset#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#http_header Daemonset#http_header}
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// Path to access on the HTTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#path Daemonset#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Name or number of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#port Daemonset#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Scheme to use for connecting to the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/daemonset#scheme Daemonset#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

