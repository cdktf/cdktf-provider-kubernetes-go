//go:build no_runtime_type_checking

package podsecuritypolicy

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSecurityPolicySpecHostPortsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSecurityPolicySpecHostPortsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecHostPortsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecHostPortsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecHostPortsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecHostPortsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSecurityPolicySpecHostPortsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

