// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datakubernetesingress

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesIngressStatusLoadBalancerIngressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressStatusLoadBalancerIngressList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressStatusLoadBalancerIngressList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusLoadBalancerIngressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusLoadBalancerIngressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressStatusLoadBalancerIngressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesIngressStatusLoadBalancerIngressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

