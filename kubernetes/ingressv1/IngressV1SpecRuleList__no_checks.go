// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ingressv1

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressV1SpecRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressV1SpecRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressV1SpecRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IngressV1SpecRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressV1SpecRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressV1SpecRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressV1SpecRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

