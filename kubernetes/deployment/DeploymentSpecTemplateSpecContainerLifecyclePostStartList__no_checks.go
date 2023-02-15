//go:build no_runtime_type_checking

package deployment

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLifecyclePostStartList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecContainerLifecyclePostStartList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLifecyclePostStartList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLifecyclePostStartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLifecyclePostStartList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecContainerLifecyclePostStartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentSpecTemplateSpecContainerLifecyclePostStartListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

