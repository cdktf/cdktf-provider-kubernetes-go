// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecTemplateSpecInitContainerEnvValueFrom struct {
	// config_map_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/stateful_set_v1#config_map_key_ref StatefulSetV1#config_map_key_ref}
	ConfigMapKeyRef *StatefulSetV1SpecTemplateSpecInitContainerEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/stateful_set_v1#field_ref StatefulSetV1#field_ref}
	FieldRef *StatefulSetV1SpecTemplateSpecInitContainerEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// resource_field_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/stateful_set_v1#resource_field_ref StatefulSetV1#resource_field_ref}
	ResourceFieldRef *StatefulSetV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// secret_key_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/stateful_set_v1#secret_key_ref StatefulSetV1#secret_key_ref}
	SecretKeyRef *StatefulSetV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

