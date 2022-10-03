//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package servicev1

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceV1SpecPortList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceV1SpecPortList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceV1SpecPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceV1SpecPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceV1SpecPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceV1SpecPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceV1SpecPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

