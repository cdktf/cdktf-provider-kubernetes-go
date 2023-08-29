// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package pod

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodSpecContainerLifecyclePreStopTcpSocketList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodSpecContainerLifecyclePreStopTcpSocketList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerLifecyclePreStopTcpSocketList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerLifecyclePreStopTcpSocketList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerLifecyclePreStopTcpSocketList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodSpecContainerLifecyclePreStopTcpSocketList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodSpecContainerLifecyclePreStopTcpSocketListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

