// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecContainerReadinessProbeExec struct {
	// Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem.
	//
	// The command is simply exec'd, it is not run inside a shell, so traditional shell instructions. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/pod#command Pod#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
}

