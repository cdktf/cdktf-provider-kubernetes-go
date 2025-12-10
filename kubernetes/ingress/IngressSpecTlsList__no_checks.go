// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ingress

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IngressSpecTlsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IngressSpecTlsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IngressSpecTlsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IngressSpecTlsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IngressSpecTlsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IngressSpecTlsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IngressSpecTlsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIngressSpecTlsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

