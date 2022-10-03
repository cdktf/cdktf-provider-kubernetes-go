//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package clusterrolebinding

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterRoleBindingSubjectList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterRoleBindingSubjectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleBindingSubjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleBindingSubjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleBindingSubjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleBindingSubjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterRoleBindingSubjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

