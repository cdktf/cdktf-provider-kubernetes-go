// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package statefulsetv1

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulSetV1SpecTemplateSpecContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

