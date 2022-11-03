//go:build no_runtime_type_checking

package ingress

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressSpecRuleHttpPathList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressSpecRuleHttpPathList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleHttpPathList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleHttpPathList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleHttpPathList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleHttpPathList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressSpecRuleHttpPathListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

