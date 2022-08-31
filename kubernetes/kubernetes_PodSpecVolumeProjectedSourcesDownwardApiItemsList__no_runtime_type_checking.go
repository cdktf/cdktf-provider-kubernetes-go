//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecVolumeProjectedSourcesDownwardApiItemsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecVolumeProjectedSourcesDownwardApiItemsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

