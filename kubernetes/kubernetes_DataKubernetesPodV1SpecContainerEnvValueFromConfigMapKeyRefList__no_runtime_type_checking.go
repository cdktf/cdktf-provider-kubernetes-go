//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesPodV1SpecContainerEnvValueFromConfigMapKeyRefList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesPodV1SpecContainerEnvValueFromConfigMapKeyRefList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodV1SpecContainerEnvValueFromConfigMapKeyRefList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodV1SpecContainerEnvValueFromConfigMapKeyRefList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodV1SpecContainerEnvValueFromConfigMapKeyRefList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesPodV1SpecContainerEnvValueFromConfigMapKeyRefListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
