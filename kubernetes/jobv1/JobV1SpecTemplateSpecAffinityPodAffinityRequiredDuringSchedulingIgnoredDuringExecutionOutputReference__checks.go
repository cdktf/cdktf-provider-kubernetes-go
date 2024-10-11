// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package jobv1

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validatePutLabelSelectorParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector:
		value := value.(*[]*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector:
		value_ := value.([]*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validatePutNamespaceSelectorParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector:
		value := value.(*[]*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector:
		value_ := value.([]*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution:
		val := val.(*JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution:
		val_ := val.(JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateSetNamespacesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReference) validateSetTopologyKeyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewJobV1SpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

