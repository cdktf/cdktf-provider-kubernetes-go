// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1


type StatefulSetV1SpecTemplateSpecInitContainerEnvValueFromResourceFieldRef struct {
	// Resource to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set_v1#resource StatefulSetV1#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set_v1#container_name StatefulSetV1#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.32.0/docs/resources/stateful_set_v1#divisor StatefulSetV1#divisor}.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

