// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageclass

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageClass) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validatePutAllowedTopologiesParameters(value *StorageClassAllowedTopologies) error {
	return nil
}

func (s *jsiiProxy_StorageClass) validatePutMetadataParameters(value *StorageClassMetadata) error {
	return nil
}

func validateStorageClass_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStorageClass_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateStorageClass_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetAllowVolumeExpansionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetMountOptionsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetParametersParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetReclaimPolicyParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetStorageProvisionerParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageClass) validateSetVolumeBindingModeParameters(val *string) error {
	return nil
}

func validateNewStorageClassParameters(scope constructs.Construct, id *string, config *StorageClassConfig) error {
	return nil
}

