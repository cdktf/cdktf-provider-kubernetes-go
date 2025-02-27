// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItems struct {
	// Path is the relative path name of the file to be created.
	//
	// Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller_v1#path ReplicationControllerV1#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller_v1#field_ref ReplicationControllerV1#field_ref}
	FieldRef *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// Mode bits to use on this file, must be a value between 0 and 0777.
	//
	// If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller_v1#mode ReplicationControllerV1#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/replication_controller_v1#resource_field_ref ReplicationControllerV1#resource_field_ref}
	ResourceFieldRef *ReplicationControllerV1SpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
}

