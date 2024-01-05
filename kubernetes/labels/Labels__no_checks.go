// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package labels

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Labels) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateImportFromParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (l *jsiiProxy_Labels) validateMoveToIdParameters(id *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (l *jsiiProxy_Labels) validatePutMetadataParameters(value *LabelsMetadata) error {
	return nil
}

func validateLabels_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateLabels_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLabels_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateLabels_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetApiVersionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetFieldManagerParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetForceParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetKindParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetLabelsParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Labels) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewLabelsParameters(scope constructs.Construct, id *string, config *LabelsConfig) error {
	return nil
}

