//go:build no_runtime_type_checking

package servicev1

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceV1StatusLoadBalancerIngressList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceV1StatusLoadBalancerIngressList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusLoadBalancerIngressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusLoadBalancerIngressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusLoadBalancerIngressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceV1StatusLoadBalancerIngressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

