// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecContainerStartupProbeGrpcList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecContainerStartupProbeGrpcList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerStartupProbeGrpcList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerStartupProbeGrpcList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerStartupProbeGrpcList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerStartupProbeGrpcList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecContainerStartupProbeGrpcListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

