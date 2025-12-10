// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package jobv1

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeaderList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeaderList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeaderList) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeaderList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeader:
		val := val.(*[]*JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeader)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeader:
		val_ := val.([]*JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeader)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeader; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeaderList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeaderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeaderList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewJobV1SpecTemplateSpecContainerStartupProbeHttpGetHttpHeaderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

