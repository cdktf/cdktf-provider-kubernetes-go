// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes


type ValidatingWebhookConfigurationV1WebhookNamespaceSelector struct {
	// match_expressions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/validating_webhook_configuration_v1#match_expressions ValidatingWebhookConfigurationV1#match_expressions}
	MatchExpressions interface{} `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// A map of {key,value} pairs.
	//
	// A single {key,value} in the matchLabels map is equivalent to an element of `match_expressions`, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/kubernetes/r/validating_webhook_configuration_v1#match_labels ValidatingWebhookConfigurationV1#match_labels}
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}
