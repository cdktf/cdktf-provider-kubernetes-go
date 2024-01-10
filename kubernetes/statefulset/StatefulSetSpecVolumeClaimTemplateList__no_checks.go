// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package statefulset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulSetSpecVolumeClaimTemplateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StatefulSetSpecVolumeClaimTemplateList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulSetSpecVolumeClaimTemplateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecVolumeClaimTemplateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecVolumeClaimTemplateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecVolumeClaimTemplateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecVolumeClaimTemplateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulSetSpecVolumeClaimTemplateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

