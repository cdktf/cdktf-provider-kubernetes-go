//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReplicationControllerV1SpecTemplateSpecVolumeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
