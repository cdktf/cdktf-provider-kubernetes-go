// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontrollerv1


type ReplicationControllerV1SpecTemplateSpecContainerEnvValueFromFieldRef struct {
	// Version of the schema the FieldPath is written in terms of, defaults to "v1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/replication_controller_v1#api_version ReplicationControllerV1#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Path of the field to select in the specified API version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/replication_controller_v1#field_path ReplicationControllerV1#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
}

