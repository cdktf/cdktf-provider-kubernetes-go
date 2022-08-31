//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterRoleBindingV1SubjectList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ClusterRoleBindingV1SubjectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleBindingV1SubjectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleBindingV1SubjectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleBindingV1SubjectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ClusterRoleBindingV1SubjectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewClusterRoleBindingV1SubjectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

