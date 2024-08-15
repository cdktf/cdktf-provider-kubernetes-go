// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1


type DaemonSetV1SpecTemplateSpecInitContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemon_set_v1#add DaemonSetV1#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemon_set_v1#drop DaemonSetV1#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

