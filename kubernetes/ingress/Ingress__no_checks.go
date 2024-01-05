// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ingress

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Ingress) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateImportFromParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateMoveToIdParameters(id *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (i *jsiiProxy_Ingress) validatePutMetadataParameters(value *IngressMetadata) error {
	return nil
}

func (i *jsiiProxy_Ingress) validatePutSpecParameters(value *IngressSpec) error {
	return nil
}

func validateIngress_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateIngress_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIngress_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateIngress_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Ingress) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ingress) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ingress) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ingress) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Ingress) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Ingress) validateSetWaitForLoadBalancerParameters(val interface{}) error {
	return nil
}

func validateNewIngressParameters(scope constructs.Construct, id *string, config *IngressConfig) error {
	return nil
}

