// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podsecuritypolicy


type PodSecurityPolicySpecAllowedFlexVolumes struct {
	// driver is the name of the Flexvolume driver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/pod_security_policy#driver PodSecurityPolicy#driver}
	Driver *string `field:"required" json:"driver" yaml:"driver"`
}

