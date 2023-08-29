// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageclass

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageClassAllowedTopologiesMatchLabelExpressionsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageClassAllowedTopologiesMatchLabelExpressionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageClassAllowedTopologiesMatchLabelExpressionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageClassAllowedTopologiesMatchLabelExpressionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageClassAllowedTopologiesMatchLabelExpressionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageClassAllowedTopologiesMatchLabelExpressionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageClassAllowedTopologiesMatchLabelExpressionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

