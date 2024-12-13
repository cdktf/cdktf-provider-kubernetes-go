// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mutatingwebhookconfigurationv1


type MutatingWebhookConfigurationV1WebhookRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/mutating_webhook_configuration_v1#api_groups MutatingWebhookConfigurationV1#api_groups}.
	ApiGroups *[]*string `field:"required" json:"apiGroups" yaml:"apiGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/mutating_webhook_configuration_v1#api_versions MutatingWebhookConfigurationV1#api_versions}.
	ApiVersions *[]*string `field:"required" json:"apiVersions" yaml:"apiVersions"`
	// Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added.
	//
	// If '*' is present, the length of the slice must be one. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/mutating_webhook_configuration_v1#operations MutatingWebhookConfigurationV1#operations}
	Operations *[]*string `field:"required" json:"operations" yaml:"operations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/mutating_webhook_configuration_v1#resources MutatingWebhookConfigurationV1#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/mutating_webhook_configuration_v1#scope MutatingWebhookConfigurationV1#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

