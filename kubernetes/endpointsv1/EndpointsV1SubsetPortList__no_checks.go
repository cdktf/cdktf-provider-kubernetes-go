// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package endpointsv1

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointsV1SubsetPortList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EndpointsV1SubsetPortList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointsV1SubsetPortList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetPortList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetPortList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetPortList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointsV1SubsetPortList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointsV1SubsetPortListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

