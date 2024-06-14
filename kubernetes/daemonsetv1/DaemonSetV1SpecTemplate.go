// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1


type DaemonSetV1SpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/daemon_set_v1#metadata DaemonSetV1#metadata}
	Metadata *DaemonSetV1SpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.31.0/docs/resources/daemon_set_v1#spec DaemonSetV1#spec}
	Spec *DaemonSetV1SpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

