package job


type JobSpecTemplateSpecInitContainerLifecyclePostStart struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#exec Job#exec}
	Exec *JobSpecTemplateSpecInitContainerLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#http_get Job#http_get}
	HttpGet *JobSpecTemplateSpecInitContainerLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job#tcp_socket Job#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

