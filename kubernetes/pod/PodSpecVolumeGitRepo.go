// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecVolumeGitRepo struct {
	// Target directory name.
	//
	// Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod#directory Pod#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Repository URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod#repository Pod#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// Commit hash for the specified revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/pod#revision Pod#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
}

