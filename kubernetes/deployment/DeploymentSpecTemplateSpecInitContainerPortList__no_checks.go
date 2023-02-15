//go:build no_runtime_type_checking

package deployment

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerPortList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecInitContainerPortList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecInitContainerPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentSpecTemplateSpecInitContainerPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

