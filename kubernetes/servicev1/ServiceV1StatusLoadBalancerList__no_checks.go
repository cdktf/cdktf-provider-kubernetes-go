// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package servicev1

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceV1StatusLoadBalancerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceV1StatusLoadBalancerList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceV1StatusLoadBalancerList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusLoadBalancerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusLoadBalancerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceV1StatusLoadBalancerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceV1StatusLoadBalancerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

