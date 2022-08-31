//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MutatingWebhookConfigurationWebhookList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MutatingWebhookConfigurationWebhookList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMutatingWebhookConfigurationWebhookListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

