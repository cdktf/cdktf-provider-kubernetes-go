// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobSpecTemplateSpecReadinessGateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecReadinessGateList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecReadinessGateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecReadinessGateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecReadinessGateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecReadinessGateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecReadinessGateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobSpecTemplateSpecReadinessGateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

