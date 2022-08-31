//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecDnsConfigOptionList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecDnsConfigOptionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecDnsConfigOptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecDnsConfigOptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecDnsConfigOptionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecDnsConfigOptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentV1SpecTemplateSpecDnsConfigOptionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

