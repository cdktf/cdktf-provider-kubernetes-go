// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecTemplateSpecInitContainerLifecyclePreStop struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/job#exec Job#exec}
	Exec *JobSpecTemplateSpecInitContainerLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/job#http_get Job#http_get}
	HttpGet *JobSpecTemplateSpecInitContainerLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/job#tcp_socket Job#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

