// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package replicationcontroller


type ReplicationControllerSpecTemplateSpecSecurityContextWindowsOptions struct {
	// GMSACredentialSpec is where the GMSA admission webhook inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/replication_controller#gmsa_credential_spec ReplicationController#gmsa_credential_spec}
	GmsaCredentialSpec *string `field:"optional" json:"gmsaCredentialSpec" yaml:"gmsaCredentialSpec"`
	// GMSACredentialSpecName is the name of the GMSA credential spec to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/replication_controller#gmsa_credential_spec_name ReplicationController#gmsa_credential_spec_name}
	GmsaCredentialSpecName *string `field:"optional" json:"gmsaCredentialSpecName" yaml:"gmsaCredentialSpecName"`
	// HostProcess determines if a container should be run as a 'Host Process' container. Default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/replication_controller#host_process ReplicationController#host_process}
	HostProcess interface{} `field:"optional" json:"hostProcess" yaml:"hostProcess"`
	// The UserName in Windows to run the entrypoint of the container process.
	//
	// Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/kubernetes/2.35.0/docs/resources/replication_controller#run_as_username ReplicationController#run_as_username}
	RunAsUsername *string `field:"optional" json:"runAsUsername" yaml:"runAsUsername"`
}

