// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecSecurityContextSysctl struct {
	// Name of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemonset#name Daemonset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Value of a property to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/daemonset#value Daemonset#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

