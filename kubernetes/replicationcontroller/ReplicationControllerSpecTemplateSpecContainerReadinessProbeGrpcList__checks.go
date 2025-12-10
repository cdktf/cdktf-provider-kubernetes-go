// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package replicationcontroller

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpcList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpcList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpcList) validateResolveParameters(context cdktf.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpcList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpc:
		val := val.(*[]*ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpc)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpc:
		val_ := val.([]*ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpc)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *[]*ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpc; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpcList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpcList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpcList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewReplicationControllerSpecTemplateSpecContainerReadinessProbeGrpcListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

