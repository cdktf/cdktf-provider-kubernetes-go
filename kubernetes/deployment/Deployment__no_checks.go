// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deployment

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Deployment) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateImportFromParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateMoveToIdParameters(id *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (d *jsiiProxy_Deployment) validatePutMetadataParameters(value *DeploymentMetadata) error {
	return nil
}

func (d *jsiiProxy_Deployment) validatePutSpecParameters(value *DeploymentSpec) error {
	return nil
}

func (d *jsiiProxy_Deployment) validatePutTimeoutsParameters(value *DeploymentTimeouts) error {
	return nil
}

func validateDeployment_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateDeployment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDeployment_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateDeployment_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Deployment) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Deployment) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Deployment) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Deployment) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Deployment) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Deployment) validateSetWaitForRolloutParameters(val interface{}) error {
	return nil
}

func validateNewDeploymentParameters(scope constructs.Construct, id *string, config *DeploymentConfig) error {
	return nil
}

