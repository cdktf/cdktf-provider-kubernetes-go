// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datakubernetesnodes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesNodesNodesMetadataList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesNodesNodesMetadataList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesNodesNodesMetadataList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNodesNodesMetadataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNodesNodesMetadataList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNodesNodesMetadataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesNodesNodesMetadataListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

