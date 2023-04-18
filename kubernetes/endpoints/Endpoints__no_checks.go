//go:build no_runtime_type_checking

package endpoints

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Endpoints) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validatePutMetadataParameters(value *EndpointsMetadata) error {
	return nil
}

func (e *jsiiProxy_Endpoints) validatePutSubsetParameters(value interface{}) error {
	return nil
}

func validateEndpoints_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEndpoints_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateEndpoints_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Endpoints) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Endpoints) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Endpoints) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Endpoints) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Endpoints) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewEndpointsParameters(scope constructs.Construct, id *string, config *EndpointsConfig) error {
	return nil
}

