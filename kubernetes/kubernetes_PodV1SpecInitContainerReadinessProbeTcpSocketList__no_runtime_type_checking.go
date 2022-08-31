//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodV1SpecInitContainerReadinessProbeTcpSocketList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodV1SpecInitContainerReadinessProbeTcpSocketList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecInitContainerReadinessProbeTcpSocketList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecInitContainerReadinessProbeTcpSocketList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecInitContainerReadinessProbeTcpSocketList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecInitContainerReadinessProbeTcpSocketList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodV1SpecInitContainerReadinessProbeTcpSocketListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

