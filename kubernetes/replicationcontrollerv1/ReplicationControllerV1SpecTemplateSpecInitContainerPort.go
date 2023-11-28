// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecInitContainerPort struct {
	// Number of port to expose on the pod's IP address.
	//
	// This must be a valid port number, 0 < x < 65536.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#container_port ReplicationControllerV1#container_port}
	ContainerPort *float64 `field:"required" json:"containerPort" yaml:"containerPort"`
	// What host IP to bind the external port to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#host_ip ReplicationControllerV1#host_ip}
	HostIp *string `field:"optional" json:"hostIp" yaml:"hostIp"`
	// Number of port to expose on the host.
	//
	// If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#host_port ReplicationControllerV1#host_port}
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// If specified, this must be an IANA_SVC_NAME and unique within the pod.
	//
	// Each named port in a pod must have a unique name. Name for the port that can be referred to by services
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#name ReplicationControllerV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Protocol for port. Must be UDP or TCP. Defaults to "TCP".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/replication_controller_v1#protocol ReplicationControllerV1#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

