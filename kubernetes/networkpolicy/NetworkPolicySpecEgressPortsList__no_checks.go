// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPolicySpecEgressPortsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicySpecEgressPortsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicySpecEgressPortsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressPortsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressPortsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressPortsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressPortsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkPolicySpecEgressPortsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

