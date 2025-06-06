// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicev1


type ServiceV1SpecSessionAffinityConfig struct {
	// client_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/service_v1#client_ip ServiceV1#client_ip}
	ClientIp *ServiceV1SpecSessionAffinityConfigClientIp `field:"optional" json:"clientIp" yaml:"clientIp"`
}

