// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rolev1

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleV1RuleList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RoleV1RuleList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RoleV1RuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RoleV1RuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleV1RuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RoleV1RuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RoleV1RuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoleV1RuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

