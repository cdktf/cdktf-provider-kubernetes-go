// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-kubernetes.provider.KubernetesProvider",
		reflect.TypeOf((*KubernetesProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificate", GoGetter: "ClientCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "clientCertificateInput", GoGetter: "ClientCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientKey", GoGetter: "ClientKey"},
			_jsii_.MemberProperty{JsiiProperty: "clientKeyInput", GoGetter: "ClientKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCaCertificate", GoGetter: "ClusterCaCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCaCertificateInput", GoGetter: "ClusterCaCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "configContext", GoGetter: "ConfigContext"},
			_jsii_.MemberProperty{JsiiProperty: "configContextAuthInfo", GoGetter: "ConfigContextAuthInfo"},
			_jsii_.MemberProperty{JsiiProperty: "configContextAuthInfoInput", GoGetter: "ConfigContextAuthInfoInput"},
			_jsii_.MemberProperty{JsiiProperty: "configContextCluster", GoGetter: "ConfigContextCluster"},
			_jsii_.MemberProperty{JsiiProperty: "configContextClusterInput", GoGetter: "ConfigContextClusterInput"},
			_jsii_.MemberProperty{JsiiProperty: "configContextInput", GoGetter: "ConfigContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "configPath", GoGetter: "ConfigPath"},
			_jsii_.MemberProperty{JsiiProperty: "configPathInput", GoGetter: "ConfigPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "configPaths", GoGetter: "ConfigPaths"},
			_jsii_.MemberProperty{JsiiProperty: "configPathsInput", GoGetter: "ConfigPathsInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "exec", GoGetter: "Exec"},
			_jsii_.MemberProperty{JsiiProperty: "execInput", GoGetter: "ExecInput"},
			_jsii_.MemberProperty{JsiiProperty: "experiments", GoGetter: "Experiments"},
			_jsii_.MemberProperty{JsiiProperty: "experimentsInput", GoGetter: "ExperimentsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "hostInput", GoGetter: "HostInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreAnnotations", GoGetter: "IgnoreAnnotations"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreAnnotationsInput", GoGetter: "IgnoreAnnotationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreLabels", GoGetter: "IgnoreLabels"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreLabelsInput", GoGetter: "IgnoreLabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "proxyUrl", GoGetter: "ProxyUrl"},
			_jsii_.MemberProperty{JsiiProperty: "proxyUrlInput", GoGetter: "ProxyUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCertificate", GoMethod: "ResetClientCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientKey", GoMethod: "ResetClientKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterCaCertificate", GoMethod: "ResetClusterCaCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigContext", GoMethod: "ResetConfigContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigContextAuthInfo", GoMethod: "ResetConfigContextAuthInfo"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigContextCluster", GoMethod: "ResetConfigContextCluster"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigPath", GoMethod: "ResetConfigPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfigPaths", GoMethod: "ResetConfigPaths"},
			_jsii_.MemberMethod{JsiiMethod: "resetExec", GoMethod: "ResetExec"},
			_jsii_.MemberMethod{JsiiMethod: "resetExperiments", GoMethod: "ResetExperiments"},
			_jsii_.MemberMethod{JsiiMethod: "resetHost", GoMethod: "ResetHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreAnnotations", GoMethod: "ResetIgnoreAnnotations"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreLabels", GoMethod: "ResetIgnoreLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetProxyUrl", GoMethod: "ResetProxyUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsServerName", GoMethod: "ResetTlsServerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsername", GoMethod: "ResetUsername"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsServerName", GoGetter: "TlsServerName"},
			_jsii_.MemberProperty{JsiiProperty: "tlsServerNameInput", GoGetter: "TlsServerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "username", GoGetter: "Username"},
			_jsii_.MemberProperty{JsiiProperty: "usernameInput", GoGetter: "UsernameInput"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-kubernetes.provider.KubernetesProviderConfig",
		reflect.TypeOf((*KubernetesProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-kubernetes.provider.KubernetesProviderExec",
		reflect.TypeOf((*KubernetesProviderExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-kubernetes.provider.KubernetesProviderExperiments",
		reflect.TypeOf((*KubernetesProviderExperiments)(nil)).Elem(),
	)
}
