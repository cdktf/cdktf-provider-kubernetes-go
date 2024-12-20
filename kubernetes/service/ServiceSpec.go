// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceSpec struct {
	// Defines if `NodePorts` will be automatically allocated for services with type `LoadBalancer`.
	//
	// It may be set to `false` if the cluster load-balancer does not rely on `NodePorts`.  If the caller requests specific `NodePorts` (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type `LoadBalancer`. Default is `true`. More info: https://kubernetes.io/docs/concepts/services-networking/service/#load-balancer-nodeport-allocation
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#allocate_load_balancer_node_ports Service#allocate_load_balancer_node_ports}
	AllocateLoadBalancerNodePorts interface{} `field:"optional" json:"allocateLoadBalancerNodePorts" yaml:"allocateLoadBalancerNodePorts"`
	// The IP address of the service.
	//
	// It is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. `None` can be specified for headless services when proxying is not required. Ignored if type is `ExternalName`. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#cluster_ip Service#cluster_ip}
	ClusterIp *string `field:"optional" json:"clusterIp" yaml:"clusterIp"`
	// List of IP addresses assigned to this service, and are usually assigned randomly.
	//
	// If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise creation of the service will fail. If this field is not specified, it will be initialized from the `clusterIP` field. If this field is specified, clients must ensure that `clusterIPs[0]` and `clusterIP` have the same value. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#cluster_ips Service#cluster_ips}
	ClusterIps *[]*string `field:"optional" json:"clusterIps" yaml:"clusterIps"`
	// A list of IP addresses for which nodes in the cluster will also accept traffic for this service.
	//
	// These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#external_ips Service#external_ips}
	ExternalIps *[]*string `field:"optional" json:"externalIps" yaml:"externalIps"`
	// The external reference that kubedns or equivalent will return as a CNAME record for this service.
	//
	// No proxying will be involved. Must be a valid DNS name and requires `type` to be `ExternalName`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#external_name Service#external_name}
	ExternalName *string `field:"optional" json:"externalName" yaml:"externalName"`
	// Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints.
	//
	// `Local` preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. `Cluster` obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading. More info: https://kubernetes.io/docs/tutorials/services/source-ip/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#external_traffic_policy Service#external_traffic_policy}
	ExternalTrafficPolicy *string `field:"optional" json:"externalTrafficPolicy" yaml:"externalTrafficPolicy"`
	// Specifies the Healthcheck NodePort for the service.
	//
	// Only effects when type is set to `LoadBalancer` and external_traffic_policy is set to `Local`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#health_check_node_port Service#health_check_node_port}
	HealthCheckNodePort *float64 `field:"optional" json:"healthCheckNodePort" yaml:"healthCheckNodePort"`
	// Specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only.
	//
	// `Cluster` routes internal traffic to a Service to all endpoints. `Local` routes traffic to node-local endpoints only, traffic is dropped if no node-local endpoints are ready. The default value is `Cluster`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#internal_traffic_policy Service#internal_traffic_policy}
	InternalTrafficPolicy *string `field:"optional" json:"internalTrafficPolicy" yaml:"internalTrafficPolicy"`
	// IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#ip_families Service#ip_families}
	IpFamilies *[]*string `field:"optional" json:"ipFamilies" yaml:"ipFamilies"`
	// IPFamilyPolicy represents the dual-stack-ness requested or required by this Service.
	//
	// If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#ip_family_policy Service#ip_family_policy}
	IpFamilyPolicy *string `field:"optional" json:"ipFamilyPolicy" yaml:"ipFamilyPolicy"`
	// The class of the load balancer implementation this Service belongs to.
	//
	// If specified, the value of this field must be a label-style identifier, with an optional prefix. This field can only be set when the Service type is `LoadBalancer`. If not set, the default load balancer implementation is used. This field can only be set when creating or updating a Service to type `LoadBalancer`. More info: https://kubernetes.io/docs/concepts/services-networking/service/#load-balancer-class
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#load_balancer_class Service#load_balancer_class}
	LoadBalancerClass *string `field:"optional" json:"loadBalancerClass" yaml:"loadBalancerClass"`
	// Only applies to `type = LoadBalancer`.
	//
	// LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying this field when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#load_balancer_ip Service#load_balancer_ip}
	LoadBalancerIp *string `field:"optional" json:"loadBalancerIp" yaml:"loadBalancerIp"`
	// If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs.
	//
	// This field will be ignored if the cloud-provider does not support the feature. More info: http://kubernetes.io/docs/user-guide/services-firewalls
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#load_balancer_source_ranges Service#load_balancer_source_ranges}
	LoadBalancerSourceRanges *[]*string `field:"optional" json:"loadBalancerSourceRanges" yaml:"loadBalancerSourceRanges"`
	// port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#port Service#port}
	Port interface{} `field:"optional" json:"port" yaml:"port"`
	// When set to true, indicates that DNS implementations must publish the `notReadyAddresses` of subsets for the Endpoints associated with the Service.
	//
	// The default value is `false`. The primary use case for setting this field is to use a StatefulSet's Headless Service to propagate `SRV` records for its Pods without respect to their readiness for purpose of peer discovery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#publish_not_ready_addresses Service#publish_not_ready_addresses}
	PublishNotReadyAddresses interface{} `field:"optional" json:"publishNotReadyAddresses" yaml:"publishNotReadyAddresses"`
	// Route service traffic to pods with label keys and values matching this selector.
	//
	// Only applies to types `ClusterIP`, `NodePort`, and `LoadBalancer`. More info: https://kubernetes.io/docs/concepts/services-networking/service/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#selector Service#selector}
	Selector *map[string]*string `field:"optional" json:"selector" yaml:"selector"`
	// Used to maintain session affinity. Supports `ClientIP` and `None`. Defaults to `None`. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#session_affinity Service#session_affinity}
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// session_affinity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#session_affinity_config Service#session_affinity_config}
	SessionAffinityConfig *ServiceSpecSessionAffinityConfig `field:"optional" json:"sessionAffinityConfig" yaml:"sessionAffinityConfig"`
	// Determines how the service is exposed.
	//
	// Defaults to `ClusterIP`. Valid options are `ExternalName`, `ClusterIP`, `NodePort`, and `LoadBalancer`. `ExternalName` maps to the specified `external_name`. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.1/docs/resources/service#type Service#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

