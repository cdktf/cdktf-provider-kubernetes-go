// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package replicationcontroller

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewReplicationControllerSpecTemplateSpecContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

