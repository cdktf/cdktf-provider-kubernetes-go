// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller


type ReplicationControllerSpecTemplateSpecInitContainerLifecyclePreStop struct {
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/replication_controller#exec ReplicationController#exec}
	Exec *ReplicationControllerSpecTemplateSpecInitContainerLifecyclePreStopExec `field:"optional" json:"exec" yaml:"exec"`
	// http_get block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/replication_controller#http_get ReplicationController#http_get}
	HttpGet *ReplicationControllerSpecTemplateSpecInitContainerLifecyclePreStopHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// tcp_socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/replication_controller#tcp_socket ReplicationController#tcp_socket}
	TcpSocket interface{} `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

