// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package manifest

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Manifest) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (m *jsiiProxy_Manifest) validatePutFieldManagerParameters(value *ManifestFieldManager) error {
	return nil
}

func (m *jsiiProxy_Manifest) validatePutTimeoutsParameters(value *ManifestTimeouts) error {
	return nil
}

func (m *jsiiProxy_Manifest) validatePutWaitParameters(value *ManifestWait) error {
	return nil
}

func (m *jsiiProxy_Manifest) validatePutWaitForParameters(value *ManifestWaitFor) error {
	return nil
}

func validateManifest_IsConstructParameters(x interface{}) error {
	return nil
}

func validateManifest_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateManifest_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Manifest) validateSetComputedFieldsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Manifest) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Manifest) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Manifest) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Manifest) validateSetManifestParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_Manifest) validateSetObjectParameters(val *map[string]interface{}) error {
	return nil
}

func (j *jsiiProxy_Manifest) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewManifestParameters(scope constructs.Construct, id *string, config *ManifestConfig) error {
	return nil
}

