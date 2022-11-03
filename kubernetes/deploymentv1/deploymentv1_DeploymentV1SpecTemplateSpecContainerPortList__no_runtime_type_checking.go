//go:build no_runtime_type_checking

package deploymentv1

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerPortList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecContainerPortList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecContainerPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentV1SpecTemplateSpecContainerPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

