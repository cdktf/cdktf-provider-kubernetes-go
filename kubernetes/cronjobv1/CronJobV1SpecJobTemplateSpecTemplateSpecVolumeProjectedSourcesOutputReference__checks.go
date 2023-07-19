//go:build !no_runtime_type_checking

package cronjobv1

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validatePutConfigMapParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesConfigMap:
		value := value.(*[]*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesConfigMap)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesConfigMap:
		value_ := value.([]*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesConfigMap)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesConfigMap; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validatePutDownwardApiParameters(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesDownwardApi) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validatePutSecretParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecret:
		value := value.(*[]*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecret)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecret:
		value_ := value.([]*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecret)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesSecret; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validatePutServiceAccountTokenParameters(value *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesServiceAccountToken) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSources:
		val := val.(*CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSources)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSources:
		val_ := val.(CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSources)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSources; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCronJobV1SpecJobTemplateSpecTemplateSpecVolumeProjectedSourcesOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

