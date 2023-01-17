package statefulsetv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/statefulsetv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference interface {
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
	Exec() StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeExecOutputReference
	ExecInput() *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeExec
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	HttpGet() StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeHttpGetOutputReference
	HttpGetInput() *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeHttpGet
	InitialDelaySeconds() *float64
	SetInitialDelaySeconds(val *float64)
	InitialDelaySecondsInput() *float64
	InternalValue() *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbe
	SetInternalValue(val *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbe)
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	SuccessThreshold() *float64
	SetSuccessThreshold(val *float64)
	SuccessThresholdInput() *float64
	TcpSocket() StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeTcpSocketList
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
	PutExec(value *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeExec)
	PutHttpGet(value *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeHttpGet)
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

// The jsii proxy struct for StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference
type jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) Exec() StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeExecOutputReference {
	var returns StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeExecOutputReference
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ExecInput() *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeExec {
	var returns *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeExec
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) HttpGet() StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeHttpGetOutputReference {
	var returns StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) HttpGetInput() *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeHttpGet {
	var returns *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) InitialDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) InitialDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) InternalValue() *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbe {
	var returns *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbe
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) SuccessThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) SuccessThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) TcpSocket() StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeTcpSocketList {
	var returns StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeTcpSocketList
	_jsii_.Get(
		j,
		"tcpSocket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) TcpSocketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpSocketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewStatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSetV1.StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference_Override(s StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSetV1.StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetInitialDelaySeconds(val *float64) {
	if err := j.validateSetInitialDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetInternalValue(val *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbe) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetSuccessThreshold(val *float64) {
	if err := j.validateSetSuccessThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successThreshold",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) PutExec(value *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeExec) {
	if err := s.validatePutExecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExec",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) PutHttpGet(value *StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeHttpGet) {
	if err := s.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHttpGet",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) PutTcpSocket(value interface{}) {
	if err := s.validatePutTcpSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTcpSocket",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ResetExec() {
	_jsii_.InvokeVoid(
		s,
		"resetExec",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		s,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ResetInitialDelaySeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialDelaySeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ResetSuccessThreshold() {
	_jsii_.InvokeVoid(
		s,
		"resetSuccessThreshold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ResetTcpSocket() {
	_jsii_.InvokeVoid(
		s,
		"resetTcpSocket",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecInitContainerLivenessProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

