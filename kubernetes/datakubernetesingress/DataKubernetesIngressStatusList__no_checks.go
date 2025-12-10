// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datakubernetesingress

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesIngressStatusList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressStatusList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesIngressStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

