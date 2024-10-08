// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package horizontalpodautoscaler

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleDownList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleDownList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleDownList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleDownList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleDownList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleDownList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleDownList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHorizontalPodAutoscalerSpecBehaviorScaleDownListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

