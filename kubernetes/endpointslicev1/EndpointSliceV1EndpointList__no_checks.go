// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package endpointslicev1

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EndpointSliceV1EndpointList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EndpointSliceV1EndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EndpointSliceV1EndpointList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EndpointSliceV1EndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EndpointSliceV1EndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EndpointSliceV1EndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EndpointSliceV1EndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEndpointSliceV1EndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

