// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package deploymentv1

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecInitContainerEnvFromList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecInitContainerEnvFromList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecInitContainerEnvFromList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecInitContainerEnvFromList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecInitContainerEnvFromList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecInitContainerEnvFromList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDeploymentV1SpecTemplateSpecInitContainerEnvFromListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

