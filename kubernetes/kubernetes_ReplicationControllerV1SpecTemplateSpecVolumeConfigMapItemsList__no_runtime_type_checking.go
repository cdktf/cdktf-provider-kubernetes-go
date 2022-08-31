//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeConfigMapItemsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeConfigMapItemsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeConfigMapItemsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeConfigMapItemsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeConfigMapItemsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecVolumeConfigMapItemsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReplicationControllerV1SpecTemplateSpecVolumeConfigMapItemsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

