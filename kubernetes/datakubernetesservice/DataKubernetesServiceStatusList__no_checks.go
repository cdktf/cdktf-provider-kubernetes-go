// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datakubernetesservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesServiceStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesServiceStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesServiceStatusList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesServiceStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

