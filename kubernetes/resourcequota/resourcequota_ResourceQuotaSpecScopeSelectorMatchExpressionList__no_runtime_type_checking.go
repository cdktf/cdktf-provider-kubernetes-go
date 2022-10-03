//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package resourcequota

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourceQuotaSpecScopeSelectorMatchExpressionList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ResourceQuotaSpecScopeSelectorMatchExpressionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ResourceQuotaSpecScopeSelectorMatchExpressionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ResourceQuotaSpecScopeSelectorMatchExpressionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ResourceQuotaSpecScopeSelectorMatchExpressionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ResourceQuotaSpecScopeSelectorMatchExpressionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewResourceQuotaSpecScopeSelectorMatchExpressionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

