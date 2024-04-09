// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecOs struct {
	// Name is the name of the operating system. The currently supported values are linux and windows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.28.0/docs/resources/daemonset#name Daemonset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

