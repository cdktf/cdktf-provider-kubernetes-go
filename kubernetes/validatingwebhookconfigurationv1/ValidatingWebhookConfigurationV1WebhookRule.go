// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validatingwebhookconfigurationv1


type ValidatingWebhookConfigurationV1WebhookRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/validating_webhook_configuration_v1#api_groups ValidatingWebhookConfigurationV1#api_groups}.
	ApiGroups *[]*string `field:"required" json:"apiGroups" yaml:"apiGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/validating_webhook_configuration_v1#api_versions ValidatingWebhookConfigurationV1#api_versions}.
	ApiVersions *[]*string `field:"required" json:"apiVersions" yaml:"apiVersions"`
	// Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added.
	//
	// If '*' is present, the length of the slice must be one. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/validating_webhook_configuration_v1#operations ValidatingWebhookConfigurationV1#operations}
	Operations *[]*string `field:"required" json:"operations" yaml:"operations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/validating_webhook_configuration_v1#resources ValidatingWebhookConfigurationV1#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.2/docs/resources/validating_webhook_configuration_v1#scope ValidatingWebhookConfigurationV1#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

