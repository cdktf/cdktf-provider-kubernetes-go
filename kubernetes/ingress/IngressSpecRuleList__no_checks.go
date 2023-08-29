// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ingress

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressSpecRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressSpecRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressSpecRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressSpecRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

