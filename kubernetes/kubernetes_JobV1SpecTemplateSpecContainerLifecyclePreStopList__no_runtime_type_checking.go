//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerLifecyclePreStopList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerLifecyclePreStopList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerLifecyclePreStopList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerLifecyclePreStopList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerLifecyclePreStopList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerLifecyclePreStopList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobV1SpecTemplateSpecContainerLifecyclePreStopListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
