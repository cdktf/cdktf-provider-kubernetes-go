//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecTopologySpreadConstraintLabelSelectorList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecTopologySpreadConstraintLabelSelectorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecTopologySpreadConstraintLabelSelectorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecTopologySpreadConstraintLabelSelectorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecTopologySpreadConstraintLabelSelectorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecTopologySpreadConstraintLabelSelectorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecTopologySpreadConstraintLabelSelectorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

