//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesPodV1SpecInitContainerPortList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesPodV1SpecInitContainerPortList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodV1SpecInitContainerPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodV1SpecInitContainerPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodV1SpecInitContainerPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesPodV1SpecInitContainerPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
