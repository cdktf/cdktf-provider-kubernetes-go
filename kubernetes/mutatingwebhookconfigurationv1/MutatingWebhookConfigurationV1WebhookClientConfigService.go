// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mutatingwebhookconfigurationv1


type MutatingWebhookConfigurationV1WebhookClientConfigService struct {
	// `name` is the name of the service. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/mutating_webhook_configuration_v1#name MutatingWebhookConfigurationV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// `namespace` is the namespace of the service. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/mutating_webhook_configuration_v1#namespace MutatingWebhookConfigurationV1#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// `path` is an optional URL path which will be sent in any request to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/mutating_webhook_configuration_v1#path MutatingWebhookConfigurationV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If specified, the port on the service that hosting webhook.
	//
	// Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.34.0/docs/resources/mutating_webhook_configuration_v1#port MutatingWebhookConfigurationV1#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

