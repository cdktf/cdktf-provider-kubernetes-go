//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ingressv1

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressV1StatusList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressV1StatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressV1StatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressV1StatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressV1StatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressV1StatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

