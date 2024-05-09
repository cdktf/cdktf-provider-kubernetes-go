// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller


type ReplicationControllerSpecTemplateSpecContainerStartupProbeHttpGet struct {
	// Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/replication_controller#host ReplicationController#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/replication_controller#http_header ReplicationController#http_header}
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// Path to access on the HTTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/replication_controller#path ReplicationController#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Name or number of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/replication_controller#port ReplicationController#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Scheme to use for connecting to the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.30.0/docs/resources/replication_controller#scheme ReplicationController#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

