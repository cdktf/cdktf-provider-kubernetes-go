package job

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v8/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v8/job/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate
	SetInternalValue(val *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate)
	Metadata() JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadataOutputReference
	MetadataInput() *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata
	Spec() JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference
	SpecInput() *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutMetadata(value *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata)
	PutSpec(value *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec)
	ResetMetadata()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference
type jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) InternalValue() *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate {
	var returns *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) Metadata() JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadataOutputReference {
	var returns JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadataOutputReference
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) MetadataInput() *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata {
	var returns *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) Spec() JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference {
	var returns JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpecOutputReference
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) SpecInput() *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec {
	var returns *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec
	_jsii_.Get(
		j,
		"specInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewJobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference {
	_init_.Initialize()

	if err := validateNewJobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewJobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference_Override(j JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.job.JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		j,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference)SetInternalValue(val *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplate) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) PutMetadata(value *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateMetadata) {
	if err := j.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putMetadata",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) PutSpec(value *JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateSpec) {
	if err := j.validatePutSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putSpec",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		j,
		"resetMetadata",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := j.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		j,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JobSpecTemplateSpecVolumeEphemeralVolumeClaimTemplateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

