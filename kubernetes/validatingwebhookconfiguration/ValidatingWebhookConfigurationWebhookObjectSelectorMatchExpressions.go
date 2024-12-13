// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validatingwebhookconfiguration


type ValidatingWebhookConfigurationWebhookObjectSelectorMatchExpressions struct {
	// The label key that the selector applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#key ValidatingWebhookConfiguration#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A key's relationship to a set of values. Valid operators ard `In`, `NotIn`, `Exists` and `DoesNotExist`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#operator ValidatingWebhookConfiguration#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// An array of string values.
	//
	// If the operator is `In` or `NotIn`, the values array must be non-empty. If the operator is `Exists` or `DoesNotExist`, the values array must be empty. This array is replaced during a strategic merge patch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/validating_webhook_configuration#values ValidatingWebhookConfiguration#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

