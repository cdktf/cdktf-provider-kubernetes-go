// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package env


type EnvEnvValueFromConfigMapKeyRef struct {
	// The key to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/env#key Env#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/env#name Env#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its key must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.36.0/docs/resources/env#optional Env#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

