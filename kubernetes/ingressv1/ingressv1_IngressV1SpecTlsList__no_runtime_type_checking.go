//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ingressv1

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressV1SpecTlsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressV1SpecTlsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressV1SpecTlsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IngressV1SpecTlsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressV1SpecTlsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressV1SpecTlsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressV1SpecTlsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

