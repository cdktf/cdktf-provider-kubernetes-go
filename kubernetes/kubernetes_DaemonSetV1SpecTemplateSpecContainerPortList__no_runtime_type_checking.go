//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecContainerPortList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecContainerPortList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecContainerPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecContainerPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecContainerPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecContainerPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDaemonSetV1SpecTemplateSpecContainerPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
