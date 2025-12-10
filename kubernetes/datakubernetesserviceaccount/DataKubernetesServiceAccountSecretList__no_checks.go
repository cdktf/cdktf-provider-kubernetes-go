// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datakubernetesserviceaccount

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesServiceAccountSecretList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesServiceAccountSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesServiceAccountSecretList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceAccountSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceAccountSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceAccountSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesServiceAccountSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

