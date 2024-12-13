// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecInitContainerLifecyclePreStopHttpGet struct {
	// Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#host PodV1#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// http_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#http_header PodV1#http_header}
	HttpHeader interface{} `field:"optional" json:"httpHeader" yaml:"httpHeader"`
	// Path to access on the HTTP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#path PodV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Name or number of the port to access on the container.
	//
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#port PodV1#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// Scheme to use for connecting to the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/pod_v1#scheme PodV1#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

