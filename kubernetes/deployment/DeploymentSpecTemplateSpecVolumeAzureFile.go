// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package deployment


type DeploymentSpecTemplateSpecVolumeAzureFile struct {
	// The name of secret that contains Azure Storage Account Name and Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment#secret_name Deployment#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Share Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment#share_name Deployment#share_name}
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// Whether to force the read-only setting in VolumeMounts. Defaults to false (read/write).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment#read_only Deployment#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// The namespace of the secret that contains Azure Storage Account Name and Key.
	//
	// For Kubernetes up to 1.18.x the default is the same as the Pod. For Kubernetes 1.19.x and later the default is "default" namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.23.0/docs/resources/deployment#secret_namespace Deployment#secret_namespace}
	SecretNamespace *string `field:"optional" json:"secretNamespace" yaml:"secretNamespace"`
}

