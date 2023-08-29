// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package jobv1

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvFromList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvFromList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvFromList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvFromList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvFromList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerEnvFromList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobV1SpecTemplateSpecContainerEnvFromListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

