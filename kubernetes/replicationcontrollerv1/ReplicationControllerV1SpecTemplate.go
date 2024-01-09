// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontrollerv1


type ReplicationControllerV1SpecTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/replication_controller_v1#metadata ReplicationControllerV1#metadata}
	Metadata *ReplicationControllerV1SpecTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/replication_controller_v1#spec ReplicationControllerV1#spec}
	Spec *ReplicationControllerV1SpecTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

