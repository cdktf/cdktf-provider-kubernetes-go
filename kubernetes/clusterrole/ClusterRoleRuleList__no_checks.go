// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package clusterrole

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterRoleRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterRoleRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterRoleRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

