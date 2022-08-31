//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt kubernetes Provider for Terraform CDK (cdktf)
package kubernetes

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecBehaviorScaleUpList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHorizontalPodAutoscalerSpecBehaviorScaleUpListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

