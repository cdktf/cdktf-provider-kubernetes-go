// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type KubernetesProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#alias KubernetesProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// PEM-encoded client certificate for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#client_certificate KubernetesProvider#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// PEM-encoded client certificate key for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#client_key KubernetesProvider#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// PEM-encoded root certificates bundle for TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#cluster_ca_certificate KubernetesProvider#cluster_ca_certificate}
	ClusterCaCertificate *string `field:"optional" json:"clusterCaCertificate" yaml:"clusterCaCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#config_context KubernetesProvider#config_context}.
	ConfigContext *string `field:"optional" json:"configContext" yaml:"configContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#config_context_auth_info KubernetesProvider#config_context_auth_info}.
	ConfigContextAuthInfo *string `field:"optional" json:"configContextAuthInfo" yaml:"configContextAuthInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#config_context_cluster KubernetesProvider#config_context_cluster}.
	ConfigContextCluster *string `field:"optional" json:"configContextCluster" yaml:"configContextCluster"`
	// Path to the kube config file. Can be set with KUBE_CONFIG_PATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#config_path KubernetesProvider#config_path}
	ConfigPath *string `field:"optional" json:"configPath" yaml:"configPath"`
	// A list of paths to kube config files. Can be set with KUBE_CONFIG_PATHS environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#config_paths KubernetesProvider#config_paths}
	ConfigPaths *[]*string `field:"optional" json:"configPaths" yaml:"configPaths"`
	// exec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#exec KubernetesProvider#exec}
	Exec interface{} `field:"optional" json:"exec" yaml:"exec"`
	// experiments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#experiments KubernetesProvider#experiments}
	Experiments interface{} `field:"optional" json:"experiments" yaml:"experiments"`
	// The hostname (in form of URI) of Kubernetes master.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#host KubernetesProvider#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// List of Kubernetes metadata annotations to ignore across all resources handled by this provider for situations where external systems are managing certain resource annotations.
	//
	// Each item is a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#ignore_annotations KubernetesProvider#ignore_annotations}
	IgnoreAnnotations *[]*string `field:"optional" json:"ignoreAnnotations" yaml:"ignoreAnnotations"`
	// List of Kubernetes metadata labels to ignore across all resources handled by this provider for situations where external systems are managing certain resource labels.
	//
	// Each item is a regular expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#ignore_labels KubernetesProvider#ignore_labels}
	IgnoreLabels *[]*string `field:"optional" json:"ignoreLabels" yaml:"ignoreLabels"`
	// Whether server should be accessed without verifying the TLS certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#insecure KubernetesProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#password KubernetesProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// URL to the proxy to be used for all API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#proxy_url KubernetesProvider#proxy_url}
	ProxyUrl *string `field:"optional" json:"proxyUrl" yaml:"proxyUrl"`
	// Server name passed to the server for SNI and is used in the client to check server certificates against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#tls_server_name KubernetesProvider#tls_server_name}
	TlsServerName *string `field:"optional" json:"tlsServerName" yaml:"tlsServerName"`
	// Token to authenticate an service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#token KubernetesProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs#username KubernetesProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

