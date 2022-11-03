//go:build no_runtime_type_checking

package statefulset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulSetSpecUpdateStrategyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulSetSpecUpdateStrategyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulSetSpecUpdateStrategyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

