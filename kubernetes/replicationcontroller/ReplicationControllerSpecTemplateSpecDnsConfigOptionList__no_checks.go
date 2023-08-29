// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package replicationcontroller

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecDnsConfigOptionList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecDnsConfigOptionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecDnsConfigOptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecDnsConfigOptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecDnsConfigOptionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecDnsConfigOptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReplicationControllerSpecTemplateSpecDnsConfigOptionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

