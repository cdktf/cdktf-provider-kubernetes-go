//go:build no_runtime_type_checking

package nodetaint

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NodeTaintTaintList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NodeTaintTaintList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NodeTaintTaintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NodeTaintTaintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NodeTaintTaintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NodeTaintTaintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNodeTaintTaintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

