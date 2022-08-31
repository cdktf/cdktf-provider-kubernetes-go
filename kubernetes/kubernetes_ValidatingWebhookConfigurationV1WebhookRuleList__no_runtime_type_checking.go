//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (v *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ValidatingWebhookConfigurationV1WebhookRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewValidatingWebhookConfigurationV1WebhookRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

