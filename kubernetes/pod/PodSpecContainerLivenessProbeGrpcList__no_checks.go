// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecContainerLivenessProbeGrpcList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecContainerLivenessProbeGrpcList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerLivenessProbeGrpcList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerLivenessProbeGrpcList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerLivenessProbeGrpcList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerLivenessProbeGrpcList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecContainerLivenessProbeGrpcListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

