//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeSecretItemsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeSecretItemsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeSecretItemsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeSecretItemsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeSecretItemsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeSecretItemsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReplicationControllerV1SpecTemplateSpecVolumeSecretItemsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

