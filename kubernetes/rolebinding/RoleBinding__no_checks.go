// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rolebinding

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleBinding) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validatePutMetadataParameters(value *RoleBindingMetadata) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validatePutRoleRefParameters(value *RoleBindingRoleRef) error {
	return nil
}

func (r *jsiiProxy_RoleBinding) validatePutSubjectParameters(value interface{}) error {
	return nil
}

func validateRoleBinding_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateRoleBinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRoleBinding_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateRoleBinding_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleBinding) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleBinding) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleBinding) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RoleBinding) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_RoleBinding) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewRoleBindingParameters(scope constructs.Construct, id *string, config *RoleBindingConfig) error {
	return nil
}

