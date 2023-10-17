// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package serviceaccount

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceAccount) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validatePutImagePullSecretParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validatePutMetadataParameters(value *ServiceAccountMetadata) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validatePutSecretParameters(value interface{}) error {
	return nil
}

func (s *jsiiProxy_ServiceAccount) validatePutTimeoutsParameters(value *ServiceAccountTimeouts) error {
	return nil
}

func validateServiceAccount_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateServiceAccount_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServiceAccount_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateServiceAccount_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceAccount) validateSetAutomountServiceAccountTokenParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceAccount) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceAccount) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceAccount) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceAccount) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ServiceAccount) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewServiceAccountParameters(scope constructs.Construct, id *string, config *ServiceAccountConfig) error {
	return nil
}

