//go:build no_runtime_type_checking

package datakubernetespod

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataKubernetesPodSpecVolumeGcePersistentDiskList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataKubernetesPodSpecVolumeGcePersistentDiskList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeGcePersistentDiskList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeGcePersistentDiskList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataKubernetesPodSpecVolumeGcePersistentDiskList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataKubernetesPodSpecVolumeGcePersistentDiskListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

