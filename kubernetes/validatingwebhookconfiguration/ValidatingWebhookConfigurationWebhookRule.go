// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validatingwebhookconfiguration


type ValidatingWebhookConfigurationWebhookRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#api_groups ValidatingWebhookConfiguration#api_groups}.
	ApiGroups *[]*string `field:"required" json:"apiGroups" yaml:"apiGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#api_versions ValidatingWebhookConfiguration#api_versions}.
	ApiVersions *[]*string `field:"required" json:"apiVersions" yaml:"apiVersions"`
	// Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added.
	//
	// If '*' is present, the length of the slice must be one. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#operations ValidatingWebhookConfiguration#operations}
	Operations *[]*string `field:"required" json:"operations" yaml:"operations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#resources ValidatingWebhookConfiguration#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#scope ValidatingWebhookConfiguration#scope}.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

