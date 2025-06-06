// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicev1


type ServiceV1SpecPort struct {
	// The port that will be exposed by this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/service_v1#port ServiceV1#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The application protocol for this port.
	//
	// This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/service_v1#app_protocol ServiceV1#app_protocol}
	AppProtocol *string `field:"optional" json:"appProtocol" yaml:"appProtocol"`
	// The name of this port within the service.
	//
	// All ports within the service must have unique names. Optional if only one ServicePort is defined on this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/service_v1#name ServiceV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on each node on which this service is exposed when `type` is `NodePort` or `LoadBalancer`.
	//
	// Usually assigned by the system. If specified, it will be allocated to the service if unused or else creation of the service will fail. Default is to auto-allocate a port if the `type` of this service requires one. More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/service_v1#node_port ServiceV1#node_port}
	NodePort *float64 `field:"optional" json:"nodePort" yaml:"nodePort"`
	// The IP protocol for this port. Supports `TCP` and `UDP`. Default is `TCP`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/service_v1#protocol ServiceV1#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Number or name of the port to access on the pods targeted by the service.
	//
	// Number must be in the range 1 to 65535. This field is ignored for services with `cluster_ip = "None"`. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/service_v1#target_port ServiceV1#target_port}
	TargetPort *string `field:"optional" json:"targetPort" yaml:"targetPort"`
}

