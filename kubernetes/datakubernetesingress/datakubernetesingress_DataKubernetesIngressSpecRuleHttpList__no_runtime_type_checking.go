//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datakubernetesingress

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesIngressSpecRuleHttpList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesIngressSpecRuleHttpList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressSpecRuleHttpList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressSpecRuleHttpList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesIngressSpecRuleHttpList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesIngressSpecRuleHttpListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
