//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulSetV1SpecVolumeClaimTemplateList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulSetV1SpecVolumeClaimTemplateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulSetV1SpecVolumeClaimTemplateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulSetV1SpecVolumeClaimTemplateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulSetV1SpecVolumeClaimTemplateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulSetV1SpecVolumeClaimTemplateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulSetV1SpecVolumeClaimTemplateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

