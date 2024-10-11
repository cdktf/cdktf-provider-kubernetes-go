// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1


type DaemonSetV1SpecTemplateSpecInitContainerEnv struct {
	// Name of the environment variable. Must be a C_IDENTIFIER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemon_set_v1#name DaemonSetV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables.
	//
	// If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemon_set_v1#value DaemonSetV1#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.33.0/docs/resources/daemon_set_v1#value_from DaemonSetV1#value_from}
	ValueFrom *DaemonSetV1SpecTemplateSpecInitContainerEnvValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

