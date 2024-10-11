// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecInitContainerVolumeDeviceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PodSpecInitContainerVolumeDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecInitContainerVolumeDeviceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerVolumeDeviceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerVolumeDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerVolumeDeviceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecInitContainerVolumeDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecInitContainerVolumeDeviceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

