//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package podsecuritypolicy

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSecurityPolicySpecRunAsGroupRangeList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSecurityPolicySpecRunAsGroupRangeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecRunAsGroupRangeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecRunAsGroupRangeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecRunAsGroupRangeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecRunAsGroupRangeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSecurityPolicySpecRunAsGroupRangeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

