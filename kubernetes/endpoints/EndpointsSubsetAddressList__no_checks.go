// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package endpoints

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsSubsetAddressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EndpointsSubsetAddressList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsSubsetAddressList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetAddressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetAddressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetAddressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsSubsetAddressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsSubsetAddressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

