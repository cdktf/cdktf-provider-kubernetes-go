// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulset


type StatefulSetSpecTemplateSpecVolumeProjectedSourcesSecret struct {
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set#items StatefulSet#items}
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set#name StatefulSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional: Specify whether the Secret or it's keys must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.38.0/docs/resources/stateful_set#optional StatefulSet#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

