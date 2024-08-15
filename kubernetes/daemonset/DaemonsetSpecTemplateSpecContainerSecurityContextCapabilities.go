// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecContainerSecurityContextCapabilities struct {
	// Added capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemonset#add Daemonset#add}
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Removed capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/daemonset#drop Daemonset#drop}
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

