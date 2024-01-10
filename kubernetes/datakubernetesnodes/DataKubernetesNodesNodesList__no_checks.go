// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datakubernetesnodes

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesNodesNodesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesNodesNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesNodesNodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNodesNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNodesNodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesNodesNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesNodesNodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

