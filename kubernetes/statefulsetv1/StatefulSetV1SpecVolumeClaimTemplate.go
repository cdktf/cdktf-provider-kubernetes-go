// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecVolumeClaimTemplate struct {
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/stateful_set_v1#metadata StatefulSetV1#metadata}
	Metadata *StatefulSetV1SpecVolumeClaimTemplateMetadata `field:"required" json:"metadata" yaml:"metadata"`
	// spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.29.0/docs/resources/stateful_set_v1#spec StatefulSetV1#spec}
	Spec *StatefulSetV1SpecVolumeClaimTemplateSpec `field:"required" json:"spec" yaml:"spec"`
}

