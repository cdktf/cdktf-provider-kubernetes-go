// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validatingwebhookconfigurationv1


type ValidatingWebhookConfigurationV1WebhookClientConfigService struct {
	// `name` is the name of the service. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/validating_webhook_configuration_v1#name ValidatingWebhookConfigurationV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// `namespace` is the namespace of the service. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/validating_webhook_configuration_v1#namespace ValidatingWebhookConfigurationV1#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// `path` is an optional URL path which will be sent in any request to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/validating_webhook_configuration_v1#path ValidatingWebhookConfigurationV1#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If specified, the port on the service that hosting webhook.
	//
	// Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.37.1/docs/resources/validating_webhook_configuration_v1#port ValidatingWebhookConfigurationV1#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

