// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package podv1


type PodV1SpecVolumeGitRepo struct {
	// Target directory name.
	//
	// Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/pod_v1#directory PodV1#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Repository URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/pod_v1#repository PodV1#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Commit hash for the specified revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/pod_v1#revision PodV1#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

