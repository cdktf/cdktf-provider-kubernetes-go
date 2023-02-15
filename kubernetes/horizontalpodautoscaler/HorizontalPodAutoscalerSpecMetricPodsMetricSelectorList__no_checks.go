//go:build no_runtime_type_checking

package horizontalpodautoscaler

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricPodsMetricSelectorList) validateGetParameters(index *float64) error {
	return nil
}

func (h *jsiiProxy_HorizontalPodAutoscalerSpecMetricPodsMetricSelectorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricPodsMetricSelectorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricPodsMetricSelectorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricPodsMetricSelectorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_HorizontalPodAutoscalerSpecMetricPodsMetricSelectorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewHorizontalPodAutoscalerSpecMetricPodsMetricSelectorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

