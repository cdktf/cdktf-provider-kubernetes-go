// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apiservicev1


type ApiServiceV1SpecService struct {
	// Name is the name of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/api_service_v1#name ApiServiceV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace is the namespace of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/api_service_v1#namespace ApiServiceV1#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// If specified, the port on the service that is hosting the service.
	//
	// Defaults to 443 for backward compatibility. Should be a valid port number (1-65535, inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/api_service_v1#port ApiServiceV1#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

