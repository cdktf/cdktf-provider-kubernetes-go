//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerLifecyclePostStartList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerLifecyclePostStartList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerLifecyclePostStartList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerLifecyclePostStartList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerLifecyclePostStartList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerLifecyclePostStartList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentSpecTemplateSpecInitContainerLifecyclePostStartListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
