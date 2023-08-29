// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package podsecuritypolicy

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSecurityPolicySpecAllowedHostPathsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSecurityPolicySpecAllowedHostPathsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecAllowedHostPathsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecAllowedHostPathsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecAllowedHostPathsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSecurityPolicySpecAllowedHostPathsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSecurityPolicySpecAllowedHostPathsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

