// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package horizontalpodautoscaler

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpPolicyList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpPolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHorizontalPodAutoscalerSpecBehaviorScaleUpPolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

