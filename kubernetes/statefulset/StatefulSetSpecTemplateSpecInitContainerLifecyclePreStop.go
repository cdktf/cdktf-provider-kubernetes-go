// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecInitContainerLifecyclePreStop struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#exec StatefulSet#exec}
	Exec *StatefulSetSpecTemplateSpecInitContainerLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#http_get StatefulSet#http_get}
	HttpGet *StatefulSetSpecTemplateSpecInitContainerLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#tcp_socket StatefulSet#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

