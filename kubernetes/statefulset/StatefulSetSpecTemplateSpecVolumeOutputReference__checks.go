//go:build !no_runtime_type_checking

package statefulset

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutAwsElasticBlockStoreParameters(value *StatefulSetSpecTemplateSpecVolumeAwsElasticBlockStore) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutAzureDiskParameters(value *StatefulSetSpecTemplateSpecVolumeAzureDisk) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutAzureFileParameters(value *StatefulSetSpecTemplateSpecVolumeAzureFile) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutCephFsParameters(value *StatefulSetSpecTemplateSpecVolumeCephFs) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutCinderParameters(value *StatefulSetSpecTemplateSpecVolumeCinder) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutConfigMapParameters(value *StatefulSetSpecTemplateSpecVolumeConfigMap) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutCsiParameters(value *StatefulSetSpecTemplateSpecVolumeCsi) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutDownwardApiParameters(value *StatefulSetSpecTemplateSpecVolumeDownwardApi) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutEmptyDirParameters(value *StatefulSetSpecTemplateSpecVolumeEmptyDir) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutFcParameters(value *StatefulSetSpecTemplateSpecVolumeFc) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutFlexVolumeParameters(value *StatefulSetSpecTemplateSpecVolumeFlexVolume) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutFlockerParameters(value *StatefulSetSpecTemplateSpecVolumeFlocker) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutGcePersistentDiskParameters(value *StatefulSetSpecTemplateSpecVolumeGcePersistentDisk) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutGitRepoParameters(value *StatefulSetSpecTemplateSpecVolumeGitRepo) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutGlusterfsParameters(value *StatefulSetSpecTemplateSpecVolumeGlusterfs) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutHostPathParameters(value *StatefulSetSpecTemplateSpecVolumeHostPath) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutIscsiParameters(value *StatefulSetSpecTemplateSpecVolumeIscsi) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutLocalParameters(value *StatefulSetSpecTemplateSpecVolumeLocal) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutNfsParameters(value *StatefulSetSpecTemplateSpecVolumeNfs) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutPersistentVolumeClaimParameters(value *StatefulSetSpecTemplateSpecVolumePersistentVolumeClaim) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutPhotonPersistentDiskParameters(value *StatefulSetSpecTemplateSpecVolumePhotonPersistentDisk) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutProjectedParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*StatefulSetSpecTemplateSpecVolumeProjected:
		value := value.(*[]*StatefulSetSpecTemplateSpecVolumeProjected)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*StatefulSetSpecTemplateSpecVolumeProjected:
		value_ := value.([]*StatefulSetSpecTemplateSpecVolumeProjected)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*StatefulSetSpecTemplateSpecVolumeProjected; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutQuobyteParameters(value *StatefulSetSpecTemplateSpecVolumeQuobyte) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutRbdParameters(value *StatefulSetSpecTemplateSpecVolumeRbd) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutSecretParameters(value *StatefulSetSpecTemplateSpecVolumeSecret) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validatePutVsphereVolumeParameters(value *StatefulSetSpecTemplateSpecVolumeVsphereVolume) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case *StatefulSetSpecTemplateSpecVolume:
		val := val.(*StatefulSetSpecTemplateSpecVolume)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case StatefulSetSpecTemplateSpecVolume:
		val_ := val.(StatefulSetSpecTemplateSpecVolume)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case cdktf.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *StatefulSetSpecTemplateSpecVolume, cdktf.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StatefulSetSpecTemplateSpecVolumeOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewStatefulSetSpecTemplateSpecVolumeOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

