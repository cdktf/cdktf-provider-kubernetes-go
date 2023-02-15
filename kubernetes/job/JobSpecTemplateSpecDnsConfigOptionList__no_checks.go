//go:build no_runtime_type_checking

package job

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JobSpecTemplateSpecDnsConfigOptionList) validateGetParameters(index *float64) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecDnsConfigOptionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecDnsConfigOptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecDnsConfigOptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecDnsConfigOptionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_JobSpecTemplateSpecDnsConfigOptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewJobSpecTemplateSpecDnsConfigOptionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

