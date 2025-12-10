// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ingressv1

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressV1StatusLoadBalancerIngressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IngressV1StatusLoadBalancerIngressList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressV1StatusLoadBalancerIngressList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressV1StatusLoadBalancerIngressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressV1StatusLoadBalancerIngressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressV1StatusLoadBalancerIngressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressV1StatusLoadBalancerIngressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

