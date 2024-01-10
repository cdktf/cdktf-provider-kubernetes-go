// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package persistentvolume

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PersistentVolumeSpecList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PersistentVolumeSpecList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PersistentVolumeSpecList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PersistentVolumeSpecList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PersistentVolumeSpecList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PersistentVolumeSpecList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PersistentVolumeSpecList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPersistentVolumeSpecListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

