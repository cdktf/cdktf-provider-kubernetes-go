//go:build no_runtime_type_checking

package namespace

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_Namespace) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (n *jsiiProxy_Namespace) validatePutMetadataParameters(value *NamespaceMetadata) error {
	return nil
}

func (n *jsiiProxy_Namespace) validatePutTimeoutsParameters(value *NamespaceTimeouts) error {
	return nil
}

func validateNamespace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNamespace_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNamespace_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Namespace) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Namespace) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Namespace) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Namespace) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Namespace) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewNamespaceParameters(scope constructs.Construct, id *string, config *NamespaceConfig) error {
	return nil
}

