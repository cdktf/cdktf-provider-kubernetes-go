// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cronjobv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v9/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v9/cronjobv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList
type jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewCronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList {
	_init_.Initialize()

	if err := validateNewCronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewCronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList_Override(c CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.cronJobV1.CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) Get(index *float64) CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CronJobV1SpecJobTemplateSpecTemplateSpecInitContainerReadinessProbeTcpSocketList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

