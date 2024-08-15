// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pod


type PodSpecContainerEnvValueFrom struct {
	// config_map_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/pod#config_map_key_ref Pod#config_map_key_ref}
	ConfigMapKeyRef *PodSpecContainerEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/pod#field_ref Pod#field_ref}
	FieldRef *PodSpecContainerEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/pod#resource_field_ref Pod#resource_field_ref}
	ResourceFieldRef *PodSpecContainerEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/pod#secret_key_ref Pod#secret_key_ref}
	SecretKeyRef *PodSpecContainerEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

