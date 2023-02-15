//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecInitContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobSpecTemplateSpecInitContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

