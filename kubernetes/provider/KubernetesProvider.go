// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs kubernetes}.
type KubernetesProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientKey() *string
	SetClientKey(val *string)
	ClientKeyInput() *string
	ClusterCaCertificate() *string
	SetClusterCaCertificate(val *string)
	ClusterCaCertificateInput() *string
	ConfigContext() *string
	SetConfigContext(val *string)
	ConfigContextAuthInfo() *string
	SetConfigContextAuthInfo(val *string)
	ConfigContextAuthInfoInput() *string
	ConfigContextCluster() *string
	SetConfigContextCluster(val *string)
	ConfigContextClusterInput() *string
	ConfigContextInput() *string
	ConfigPath() *string
	SetConfigPath(val *string)
	ConfigPathInput() *string
	ConfigPaths() *[]*string
	SetConfigPaths(val *[]*string)
	ConfigPathsInput() *[]*string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Exec() interface{}
	SetExec(val interface{})
	ExecInput() interface{}
	Experiments() interface{}
	SetExperiments(val interface{})
	ExperimentsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Host() *string
	SetHost(val *string)
	HostInput() *string
	IgnoreAnnotations() *[]*string
	SetIgnoreAnnotations(val *[]*string)
	IgnoreAnnotationsInput() *[]*string
	IgnoreLabels() *[]*string
	SetIgnoreLabels(val *[]*string)
	IgnoreLabelsInput() *[]*string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	ProxyUrl() *string
	SetProxyUrl(val *string)
	ProxyUrlInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TlsServerName() *string
	SetTlsServerName(val *string)
	TlsServerNameInput() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetClientCertificate()
	ResetClientKey()
	ResetClusterCaCertificate()
	ResetConfigContext()
	ResetConfigContextAuthInfo()
	ResetConfigContextCluster()
	ResetConfigPath()
	ResetConfigPaths()
	ResetExec()
	ResetExperiments()
	ResetHost()
	ResetIgnoreAnnotations()
	ResetIgnoreLabels()
	ResetInsecure()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetProxyUrl()
	ResetTlsServerName()
	ResetToken()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KubernetesProvider
type jsiiProxy_KubernetesProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_KubernetesProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ClientKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ClientKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ClusterCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ClusterCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigContextAuthInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configContextAuthInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigContextAuthInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configContextAuthInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigContextCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configContextCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigContextClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configContextClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConfigPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Exec() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ExecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"execInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Experiments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experiments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ExperimentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) IgnoreAnnotations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) IgnoreAnnotationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreAnnotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) IgnoreLabels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) IgnoreLabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoreLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ProxyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) ProxyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) TlsServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) TlsServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesProvider) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs kubernetes} Resource.
func NewKubernetesProvider(scope constructs.Construct, id *string, config *KubernetesProviderConfig) KubernetesProvider {
	_init_.Initialize()

	if err := validateNewKubernetesProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesProvider{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.provider.KubernetesProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.25.1/docs kubernetes} Resource.
func NewKubernetesProvider_Override(k KubernetesProvider, scope constructs.Construct, id *string, config *KubernetesProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.provider.KubernetesProvider",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetClientCertificate(val *string) {
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetClientKey(val *string) {
	_jsii_.Set(
		j,
		"clientKey",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetClusterCaCertificate(val *string) {
	_jsii_.Set(
		j,
		"clusterCaCertificate",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetConfigContext(val *string) {
	_jsii_.Set(
		j,
		"configContext",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetConfigContextAuthInfo(val *string) {
	_jsii_.Set(
		j,
		"configContextAuthInfo",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetConfigContextCluster(val *string) {
	_jsii_.Set(
		j,
		"configContextCluster",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetConfigPath(val *string) {
	_jsii_.Set(
		j,
		"configPath",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetConfigPaths(val *[]*string) {
	_jsii_.Set(
		j,
		"configPaths",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetExec(val interface{}) {
	if err := j.validateSetExecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exec",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetExperiments(val interface{}) {
	if err := j.validateSetExperimentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"experiments",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetHost(val *string) {
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetIgnoreAnnotations(val *[]*string) {
	_jsii_.Set(
		j,
		"ignoreAnnotations",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetIgnoreLabels(val *[]*string) {
	_jsii_.Set(
		j,
		"ignoreLabels",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetProxyUrl(val *string) {
	_jsii_.Set(
		j,
		"proxyUrl",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetTlsServerName(val *string) {
	_jsii_.Set(
		j,
		"tlsServerName",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_KubernetesProvider)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTF code for importing a KubernetesProvider resource upon running "cdktf plan <stack-name>".
func KubernetesProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKubernetesProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-kubernetes.provider.KubernetesProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func KubernetesProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-kubernetes.provider.KubernetesProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-kubernetes.provider.KubernetesProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-kubernetes.provider.KubernetesProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-kubernetes.provider.KubernetesProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesProvider) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesProvider) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		k,
		"resetAlias",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		k,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetClientKey() {
	_jsii_.InvokeVoid(
		k,
		"resetClientKey",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetClusterCaCertificate() {
	_jsii_.InvokeVoid(
		k,
		"resetClusterCaCertificate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetConfigContext() {
	_jsii_.InvokeVoid(
		k,
		"resetConfigContext",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetConfigContextAuthInfo() {
	_jsii_.InvokeVoid(
		k,
		"resetConfigContextAuthInfo",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetConfigContextCluster() {
	_jsii_.InvokeVoid(
		k,
		"resetConfigContextCluster",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetConfigPath() {
	_jsii_.InvokeVoid(
		k,
		"resetConfigPath",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetConfigPaths() {
	_jsii_.InvokeVoid(
		k,
		"resetConfigPaths",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetExec() {
	_jsii_.InvokeVoid(
		k,
		"resetExec",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetExperiments() {
	_jsii_.InvokeVoid(
		k,
		"resetExperiments",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetHost() {
	_jsii_.InvokeVoid(
		k,
		"resetHost",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetIgnoreAnnotations() {
	_jsii_.InvokeVoid(
		k,
		"resetIgnoreAnnotations",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetIgnoreLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetIgnoreLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		k,
		"resetInsecure",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetPassword() {
	_jsii_.InvokeVoid(
		k,
		"resetPassword",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetProxyUrl() {
	_jsii_.InvokeVoid(
		k,
		"resetProxyUrl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetTlsServerName() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsServerName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetToken() {
	_jsii_.InvokeVoid(
		k,
		"resetToken",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) ResetUsername() {
	_jsii_.InvokeVoid(
		k,
		"resetUsername",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

