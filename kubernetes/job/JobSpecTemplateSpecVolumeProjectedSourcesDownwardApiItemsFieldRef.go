// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package job


type JobSpecTemplateSpecVolumeProjectedSourcesDownwardApiItemsFieldRef struct {
	// Version of the schema the FieldPath is written in terms of, defaults to 'v1'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/job#api_version Job#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Path of the field to select in the specified API version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.24.0/docs/resources/job#field_path Job#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
}

