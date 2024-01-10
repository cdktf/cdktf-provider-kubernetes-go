// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package podv1

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodV1SpecReadinessGateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PodV1SpecReadinessGateList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PodV1SpecReadinessGateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecReadinessGateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecReadinessGateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecReadinessGateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PodV1SpecReadinessGateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPodV1SpecReadinessGateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

