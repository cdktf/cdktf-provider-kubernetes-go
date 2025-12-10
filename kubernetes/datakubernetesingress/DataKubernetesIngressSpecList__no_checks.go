// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datakubernetesingress

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesIngressSpecList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressSpecList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressSpecList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressSpecList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressSpecList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressSpecList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesIngressSpecListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

