// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package jobv1


type JobV1SpecTemplateSpecVolumeProjected struct {
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/job_v1#sources JobV1#sources}
	Sources interface{} `field:"required" json:"sources" yaml:"sources"`
	// Optional: mode bits to use on created files by default.
	//
	// Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/job_v1#default_mode JobV1#default_mode}
	DefaultMode *string `field:"optional" json:"defaultMode" yaml:"defaultMode"`
}

