package daemonsetv1


type DaemonSetV1SpecTemplateSpecContainerLifecyclePostStart struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#exec DaemonSetV1#exec}
	Exec *DaemonSetV1SpecTemplateSpecContainerLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#http_get DaemonSetV1#http_get}
	HttpGet *DaemonSetV1SpecTemplateSpecContainerLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/daemon_set_v1#tcp_socket DaemonSetV1#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

