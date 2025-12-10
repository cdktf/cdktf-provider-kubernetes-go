// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datakubernetesservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesServiceSpecSessionAffinityConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesServiceSpecSessionAffinityConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesServiceSpecSessionAffinityConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceSpecSessionAffinityConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceSpecSessionAffinityConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesServiceSpecSessionAffinityConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesServiceSpecSessionAffinityConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

