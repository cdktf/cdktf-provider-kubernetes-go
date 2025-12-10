// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package endpointsv1

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsV1SubsetAddressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EndpointsV1SubsetAddressList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsV1SubsetAddressList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetAddressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetAddressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetAddressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetAddressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsV1SubsetAddressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

