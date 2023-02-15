//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecInitContainerLifecyclePreStopList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecInitContainerLifecyclePreStopList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePreStopList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePreStopList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePreStopList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePreStopList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecInitContainerLifecyclePreStopListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

