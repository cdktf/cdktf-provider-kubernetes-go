//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecInitContainerLifecyclePostStartList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecInitContainerLifecyclePostStartList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePostStartList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePostStartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePostStartList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePostStartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecInitContainerLifecyclePostStartListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

