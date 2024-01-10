// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPolicySpecIngressFromList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicySpecIngressFromList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecIngressFromList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkPolicySpecIngressFromListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

