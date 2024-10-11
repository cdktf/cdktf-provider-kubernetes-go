// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecTemplateSpecVolumeDownwardApiItems struct {
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#field_ref Job#field_ref}
	FieldRef *JobSpecTemplateSpecVolumeDownwardApiItemsFieldRef `field:"required" json:"fieldRef" yaml:"fieldRef"`
	// Path is the relative path name of the file to be created.
	//
	// Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#path Job#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Optional: mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#mode Job#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/job#resource_field_ref Job#resource_field_ref}
	ResourceFieldRef *JobSpecTemplateSpecVolumeDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

