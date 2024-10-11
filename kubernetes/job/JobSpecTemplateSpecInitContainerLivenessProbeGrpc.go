// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecTemplateSpecInitContainerLivenessProbeGrpc struct {
	// Number of the port to access on the container. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#port Job#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#service Job#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

