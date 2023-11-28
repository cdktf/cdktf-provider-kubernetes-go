// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecVolumeLocal struct {
	// Path of the directory on the host. More info: https://kubernetes.io/docs/concepts/storage/volumes#local.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/daemonset#path Daemonset#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

