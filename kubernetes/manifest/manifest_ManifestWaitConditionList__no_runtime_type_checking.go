//go:build no_runtime_type_checking

package manifest

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManifestWaitConditionList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_ManifestWaitConditionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ManifestWaitConditionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ManifestWaitConditionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ManifestWaitConditionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ManifestWaitConditionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewManifestWaitConditionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

