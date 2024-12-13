// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecVolumeFlocker struct {
	// Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/deployment#dataset_name Deployment#dataset_name}
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// UUID of the dataset. This is unique identifier of a Flocker dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/deployment#dataset_uuid Deployment#dataset_uuid}
	DatasetUuid *string `field:"optional" json:"datasetUuid" yaml:"datasetUuid"`
}

