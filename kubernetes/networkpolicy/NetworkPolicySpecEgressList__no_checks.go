// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPolicySpecEgressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicySpecEgressList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicySpecEgressList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkPolicySpecEgressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

