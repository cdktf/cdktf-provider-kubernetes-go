//go:build no_runtime_type_checking

package daemonset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DaemonsetSpecTemplateSpecInitContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecInitContainerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecInitContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecInitContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecInitContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecInitContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDaemonsetSpecTemplateSpecInitContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

