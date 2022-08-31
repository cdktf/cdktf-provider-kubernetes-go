//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecSeLinuxSeLinuxOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSecurityPolicyV1Beta1SpecSeLinuxSeLinuxOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecSeLinuxSeLinuxOptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecSeLinuxSeLinuxOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecSeLinuxSeLinuxOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicyV1Beta1SpecSeLinuxSeLinuxOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSecurityPolicyV1Beta1SpecSeLinuxSeLinuxOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

