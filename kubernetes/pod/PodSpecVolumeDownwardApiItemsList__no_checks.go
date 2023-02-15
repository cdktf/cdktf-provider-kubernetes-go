//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecVolumeDownwardApiItemsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecVolumeDownwardApiItemsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeDownwardApiItemsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeDownwardApiItemsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeDownwardApiItemsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeDownwardApiItemsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecVolumeDownwardApiItemsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

