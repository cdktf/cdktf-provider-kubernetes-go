// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package statefulset

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StatefulSetSpecUpdateStrategyRollingUpdateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStatefulSetSpecUpdateStrategyRollingUpdateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

