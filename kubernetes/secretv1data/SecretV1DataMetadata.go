// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secretv1data


type SecretV1DataMetadata struct {
	// The name of the Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/secret_v1_data#name SecretV1Data#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of the Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/secret_v1_data#namespace SecretV1Data#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

