// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPolicySpecEgressToList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicySpecEgressToList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicySpecEgressToList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressToList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressToList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressToList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicySpecEgressToList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkPolicySpecEgressToListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

