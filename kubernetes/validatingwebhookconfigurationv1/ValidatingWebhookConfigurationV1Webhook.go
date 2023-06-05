package validatingwebhookconfigurationv1


type ValidatingWebhookConfigurationV1Webhook struct {
	// client_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#client_config ValidatingWebhookConfigurationV1#client_config}
	ClientConfig *ValidatingWebhookConfigurationV1WebhookClientConfig `field:"required" json:"clientConfig" yaml:"clientConfig"`
	// The name of the admission webhook.
	//
	// Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where "imagepolicy" is the name of the webhook, and kubernetes.io is the name of the organization. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#name ValidatingWebhookConfigurationV1#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// AdmissionReviewVersions is an ordered list of preferred `AdmissionReview` versions the Webhook expects.
	//
	// API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail and be subject to the failure policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#admission_review_versions ValidatingWebhookConfigurationV1#admission_review_versions}
	AdmissionReviewVersions *[]*string `field:"optional" json:"admissionReviewVersions" yaml:"admissionReviewVersions"`
	// FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail.
	//
	// Defaults to Fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#failure_policy ValidatingWebhookConfigurationV1#failure_policy}
	FailurePolicy *string `field:"optional" json:"failurePolicy" yaml:"failurePolicy"`
	// matchPolicy defines how the "rules" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent".
	//
	// - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.
	//
	// - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook.
	//
	// Defaults to "Equivalent"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#match_policy ValidatingWebhookConfigurationV1#match_policy}
	MatchPolicy *string `field:"optional" json:"matchPolicy" yaml:"matchPolicy"`
	// namespace_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#namespace_selector ValidatingWebhookConfigurationV1#namespace_selector}
	NamespaceSelector *ValidatingWebhookConfigurationV1WebhookNamespaceSelector `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
	// object_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#object_selector ValidatingWebhookConfigurationV1#object_selector}
	ObjectSelector *ValidatingWebhookConfigurationV1WebhookObjectSelector `field:"optional" json:"objectSelector" yaml:"objectSelector"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#rule ValidatingWebhookConfigurationV1#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// SideEffects states whether this webhook has side effects.
	//
	// Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown). Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission chain and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects == Unknown or Some.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#side_effects ValidatingWebhookConfigurationV1#side_effects}
	SideEffects *string `field:"optional" json:"sideEffects" yaml:"sideEffects"`
	// TimeoutSeconds specifies the timeout for this webhook.
	//
	// After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.21.1/docs/resources/validating_webhook_configuration_v1#timeout_seconds ValidatingWebhookConfigurationV1#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

