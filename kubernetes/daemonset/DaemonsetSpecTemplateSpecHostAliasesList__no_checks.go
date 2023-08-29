// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package daemonset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DaemonsetSpecTemplateSpecHostAliasesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DaemonsetSpecTemplateSpecHostAliasesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecHostAliasesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecHostAliasesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecHostAliasesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DaemonsetSpecTemplateSpecHostAliasesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDaemonsetSpecTemplateSpecHostAliasesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

