// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Job) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateImportFromParameters(id *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (j *jsiiProxy_Job) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (j *jsiiProxy_Job) validatePutMetadataParameters(value *JobMetadata) error {
	return nil
}

func (j *jsiiProxy_Job) validatePutSpecParameters(value *JobSpec) error {
	return nil
}

func (j *jsiiProxy_Job) validatePutTimeoutsParameters(value *JobTimeouts) error {
	return nil
}

func validateJob_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateJob_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateJob_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Job) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Job) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Job) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Job) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Job) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Job) validateSetWaitForCompletionParameters(val interface{}) error {
	return nil
}

func validateNewJobParameters(scope constructs.Construct, id *string, config *JobConfig) error {
	return nil
}

