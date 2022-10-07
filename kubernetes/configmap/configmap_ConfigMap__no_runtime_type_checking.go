//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package configmap

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigMap) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validatePutMetadataParameters(value *ConfigMapMetadata) error {
	return nil
}

func validateConfigMap_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigMap) validateSetBinaryDataParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_ConfigMap) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigMap) validateSetDataParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_ConfigMap) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigMap) validateSetImmutableParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigMap) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ConfigMap) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewConfigMapParameters(scope constructs.Construct, id *string, config *ConfigMapConfig) error {
	return nil
}

