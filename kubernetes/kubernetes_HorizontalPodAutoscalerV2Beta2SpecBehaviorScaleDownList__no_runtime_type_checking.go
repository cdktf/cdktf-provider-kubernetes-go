//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HorizontalPodAutoscalerV2Beta2SpecBehaviorScaleDownList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HorizontalPodAutoscalerV2Beta2SpecBehaviorScaleDownList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2Beta2SpecBehaviorScaleDownList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2Beta2SpecBehaviorScaleDownList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2Beta2SpecBehaviorScaleDownList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerV2Beta2SpecBehaviorScaleDownList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHorizontalPodAutoscalerV2Beta2SpecBehaviorScaleDownListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
