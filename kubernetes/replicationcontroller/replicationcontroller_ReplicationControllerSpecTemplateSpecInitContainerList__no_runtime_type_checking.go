//go:build no_runtime_type_checking

package replicationcontroller

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecInitContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecInitContainerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecInitContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecInitContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecInitContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecInitContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReplicationControllerSpecTemplateSpecInitContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

