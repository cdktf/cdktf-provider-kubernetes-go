package jobv1


type JobV1SpecTemplateSpecContainerLifecyclePostStart struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#exec JobV1#exec}
	Exec *JobV1SpecTemplateSpecContainerLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#http_get JobV1#http_get}
	HttpGet *JobV1SpecTemplateSpecContainerLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.20.0/docs/resources/job_v1#tcp_socket JobV1#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

