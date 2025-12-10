// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rolebinding

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleBindingSubjectList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RoleBindingSubjectList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RoleBindingSubjectList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RoleBindingSubjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleBindingSubjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RoleBindingSubjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RoleBindingSubjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoleBindingSubjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

