//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecInitContainerLifecyclePreStopTcpSocketList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecInitContainerLifecyclePreStopTcpSocketList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePreStopTcpSocketList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePreStopTcpSocketList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePreStopTcpSocketList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerLifecyclePreStopTcpSocketList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecInitContainerLifecyclePreStopTcpSocketListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
