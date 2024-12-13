// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonset


type DaemonsetSpecTemplateSpecVolumeSecretItems struct {
	// The key to project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/daemonset#key Daemonset#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/daemonset#mode Daemonset#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The relative path of the file to map the key to.
	//
	// May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/daemonset#path Daemonset#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

