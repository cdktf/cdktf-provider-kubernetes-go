// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecContainerEnvFrom struct {
	// config_map_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#config_map_ref StatefulSet#config_map_ref}
	ConfigMapRef *StatefulSetSpecTemplateSpecContainerEnvFromConfigMapRef `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// An optional identifer to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#prefix StatefulSet#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/stateful_set#secret_ref StatefulSet#secret_ref}
	SecretRef *StatefulSetSpecTemplateSpecContainerEnvFromSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

