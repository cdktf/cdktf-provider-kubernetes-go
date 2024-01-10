// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package podv1

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodV1SpecInitContainerEnvList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PodV1SpecInitContainerEnvList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodV1SpecInitContainerEnvList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecInitContainerEnvList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecInitContainerEnvList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecInitContainerEnvList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecInitContainerEnvList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodV1SpecInitContainerEnvListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

