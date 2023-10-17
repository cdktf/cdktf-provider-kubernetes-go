// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package secretv1

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretV1) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validatePutMetadataParameters(value *SecretV1Metadata) error {
	return nil
}

func (s *jsiiProxy_SecretV1) validatePutTimeoutsParameters(value *SecretV1Timeouts) error {
	return nil
}

func validateSecretV1_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateSecretV1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecretV1_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateSecretV1_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetBinaryDataParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetDataParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetImmutableParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecretV1) validateSetWaitForServiceAccountTokenParameters(val interface{}) error {
	return nil
}

func validateNewSecretV1Parameters(scope constructs.Construct, id *string, config *SecretV1Config) error {
	return nil
}

