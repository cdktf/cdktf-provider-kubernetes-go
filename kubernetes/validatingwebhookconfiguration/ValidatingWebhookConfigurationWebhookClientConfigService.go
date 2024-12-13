// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validatingwebhookconfiguration


type ValidatingWebhookConfigurationWebhookClientConfigService struct {
	// `name` is the name of the service. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#name ValidatingWebhookConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// `namespace` is the namespace of the service. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#namespace ValidatingWebhookConfiguration#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// `path` is an optional URL path which will be sent in any request to this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#path ValidatingWebhookConfiguration#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// If specified, the port on the service that hosting webhook.
	//
	// Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#port ValidatingWebhookConfiguration#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

