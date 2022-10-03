//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rolebindingv1

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RoleBindingV1SubjectList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RoleBindingV1SubjectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RoleBindingV1SubjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RoleBindingV1SubjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RoleBindingV1SubjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RoleBindingV1SubjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoleBindingV1SubjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

