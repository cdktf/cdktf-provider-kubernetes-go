//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodDisruptionBudgetSpecSelectorMatchExpressionsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodDisruptionBudgetSpecSelectorMatchExpressionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodDisruptionBudgetSpecSelectorMatchExpressionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodDisruptionBudgetSpecSelectorMatchExpressionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodDisruptionBudgetSpecSelectorMatchExpressionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodDisruptionBudgetSpecSelectorMatchExpressionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodDisruptionBudgetSpecSelectorMatchExpressionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

