//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package endpointsv1

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsV1SubsetNotReadyAddressList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsV1SubsetNotReadyAddressList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetNotReadyAddressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetNotReadyAddressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetNotReadyAddressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetNotReadyAddressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsV1SubsetNotReadyAddressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

