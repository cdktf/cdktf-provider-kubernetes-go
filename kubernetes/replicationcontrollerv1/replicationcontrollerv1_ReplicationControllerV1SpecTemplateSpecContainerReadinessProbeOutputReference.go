package replicationcontrollerv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/replicationcontrollerv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference interface {
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
	Exec() ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeExecOutputReference
	ExecInput() *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeExec
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	HttpGet() ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeHttpGetOutputReference
	HttpGetInput() *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeHttpGet
	InitialDelaySeconds() *float64
	SetInitialDelaySeconds(val *float64)
	InitialDelaySecondsInput() *float64
	InternalValue() *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe
	SetInternalValue(val *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe)
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	SuccessThreshold() *float64
	SetSuccessThreshold(val *float64)
	SuccessThresholdInput() *float64
	TcpSocket() ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeTcpSocketList
	TcpSocketInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutExec(value *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeExec)
	PutHttpGet(value *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeHttpGet)
	PutTcpSocket(value interface{})
	ResetExec()
	ResetFailureThreshold()
	ResetHttpGet()
	ResetInitialDelaySeconds()
	ResetPeriodSeconds()
	ResetSuccessThreshold()
	ResetTcpSocket()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference
type jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) Exec() ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeExecOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ExecInput() *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeExec {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) HttpGet() ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeHttpGetOutputReference {
	var returns ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) HttpGetInput() *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeHttpGet {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) InitialDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) InitialDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) InternalValue() *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe {
	var returns *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) SuccessThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) SuccessThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) TcpSocket() ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeTcpSocketList {
	var returns ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeTcpSocketList
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) TcpSocketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference {
	_init_.Initialize()

	if err := validateNewReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference_Override(r ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.replicationControllerV1.ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetInitialDelaySeconds(val *float64) {
	if err := j.validateSetInitialDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetInternalValue(val *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbe) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetSuccessThreshold(val *float64) {
	if err := j.validateSetSuccessThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successThreshold",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) PutExec(value *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeExec) {
	if err := r.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putExec",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) PutHttpGet(value *ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeHttpGet) {
	if err := r.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putHttpGet",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) PutTcpSocket(value interface{}) {
	if err := r.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTcpSocket",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		r,
		"resetExec",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		r,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ResetInitialDelaySeconds() {
	_jsii_.InvokeVoid(
		r,
		"resetInitialDelaySeconds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		r,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ResetSuccessThreshold() {
	_jsii_.InvokeVoid(
		r,
		"resetSuccessThreshold",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		r,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ReplicationControllerV1SpecTemplateSpecContainerReadinessProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

