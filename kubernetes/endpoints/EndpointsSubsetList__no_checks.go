// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package endpoints

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsSubsetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EndpointsSubsetList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsSubsetList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsSubsetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

