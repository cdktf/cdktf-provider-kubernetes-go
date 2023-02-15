//go:build no_runtime_type_checking

package podv1

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodV1SpecContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodV1SpecContainerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodV1SpecContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

