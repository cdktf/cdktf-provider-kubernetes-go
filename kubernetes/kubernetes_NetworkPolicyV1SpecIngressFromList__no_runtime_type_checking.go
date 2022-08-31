//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicyV1SpecIngressFromList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkPolicyV1SpecIngressFromList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkPolicyV1SpecIngressFromListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

