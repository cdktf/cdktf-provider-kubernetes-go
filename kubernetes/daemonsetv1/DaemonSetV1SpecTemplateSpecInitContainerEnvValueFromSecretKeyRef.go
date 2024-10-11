// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1


type DaemonSetV1SpecTemplateSpecInitContainerEnvValueFromSecretKeyRef struct {
	// The key of the secret to select from. Must be a valid secret key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemon_set_v1#key DaemonSetV1#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemon_set_v1#name DaemonSetV1#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the Secret or its key must be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemon_set_v1#optional DaemonSetV1#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

