//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerStartupProbeTcpSocketList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecContainerStartupProbeTcpSocketList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerStartupProbeTcpSocketList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerStartupProbeTcpSocketList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerStartupProbeTcpSocketList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecContainerStartupProbeTcpSocketList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulSetSpecTemplateSpecContainerStartupProbeTcpSocketListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
