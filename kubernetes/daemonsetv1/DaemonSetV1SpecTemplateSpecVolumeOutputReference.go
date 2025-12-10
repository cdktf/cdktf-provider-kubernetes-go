// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package daemonsetv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v12/daemonsetv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DaemonSetV1SpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() DaemonSetV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *DaemonSetV1SpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() DaemonSetV1SpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *DaemonSetV1SpecTemplateSpecVolumeAzureDisk
	AzureFile() DaemonSetV1SpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *DaemonSetV1SpecTemplateSpecVolumeAzureFile
	CephFs() DaemonSetV1SpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *DaemonSetV1SpecTemplateSpecVolumeCephFs
	Cinder() DaemonSetV1SpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *DaemonSetV1SpecTemplateSpecVolumeCinder
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
	ConfigMap() DaemonSetV1SpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *DaemonSetV1SpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() DaemonSetV1SpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *DaemonSetV1SpecTemplateSpecVolumeCsi
	DownwardApi() DaemonSetV1SpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *DaemonSetV1SpecTemplateSpecVolumeDownwardApi
	EmptyDir() DaemonSetV1SpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *DaemonSetV1SpecTemplateSpecVolumeEmptyDir
	Ephemeral() DaemonSetV1SpecTemplateSpecVolumeEphemeralOutputReference
	EphemeralInput() *DaemonSetV1SpecTemplateSpecVolumeEphemeral
	Fc() DaemonSetV1SpecTemplateSpecVolumeFcOutputReference
	FcInput() *DaemonSetV1SpecTemplateSpecVolumeFc
	FlexVolume() DaemonSetV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *DaemonSetV1SpecTemplateSpecVolumeFlexVolume
	Flocker() DaemonSetV1SpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *DaemonSetV1SpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() DaemonSetV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *DaemonSetV1SpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *DaemonSetV1SpecTemplateSpecVolumeGitRepo
	Glusterfs() DaemonSetV1SpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *DaemonSetV1SpecTemplateSpecVolumeGlusterfs
	HostPath() DaemonSetV1SpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *DaemonSetV1SpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() DaemonSetV1SpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *DaemonSetV1SpecTemplateSpecVolumeIscsi
	Local() DaemonSetV1SpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *DaemonSetV1SpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() DaemonSetV1SpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *DaemonSetV1SpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() DaemonSetV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *DaemonSetV1SpecTemplateSpecVolumePhotonPersistentDisk
	Projected() DaemonSetV1SpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() DaemonSetV1SpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *DaemonSetV1SpecTemplateSpecVolumeQuobyte
	Rbd() DaemonSetV1SpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *DaemonSetV1SpecTemplateSpecVolumeRbd
	Secret() DaemonSetV1SpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *DaemonSetV1SpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() DaemonSetV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *DaemonSetV1SpecTemplateSpecVolumeVsphereVolume
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAwsElasticBlockStore(value *DaemonSetV1SpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *DaemonSetV1SpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *DaemonSetV1SpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *DaemonSetV1SpecTemplateSpecVolumeCephFs)
	PutCinder(value *DaemonSetV1SpecTemplateSpecVolumeCinder)
	PutConfigMap(value *DaemonSetV1SpecTemplateSpecVolumeConfigMap)
	PutCsi(value *DaemonSetV1SpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *DaemonSetV1SpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *DaemonSetV1SpecTemplateSpecVolumeEmptyDir)
	PutEphemeral(value *DaemonSetV1SpecTemplateSpecVolumeEphemeral)
	PutFc(value *DaemonSetV1SpecTemplateSpecVolumeFc)
	PutFlexVolume(value *DaemonSetV1SpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *DaemonSetV1SpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *DaemonSetV1SpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *DaemonSetV1SpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *DaemonSetV1SpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *DaemonSetV1SpecTemplateSpecVolumeHostPath)
	PutIscsi(value *DaemonSetV1SpecTemplateSpecVolumeIscsi)
	PutLocal(value *DaemonSetV1SpecTemplateSpecVolumeLocal)
	PutNfs(value *DaemonSetV1SpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *DaemonSetV1SpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *DaemonSetV1SpecTemplateSpecVolumeQuobyte)
	PutRbd(value *DaemonSetV1SpecTemplateSpecVolumeRbd)
	PutSecret(value *DaemonSetV1SpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *DaemonSetV1SpecTemplateSpecVolumeVsphereVolume)
	ResetAwsElasticBlockStore()
	ResetAzureDisk()
	ResetAzureFile()
	ResetCephFs()
	ResetCinder()
	ResetConfigMap()
	ResetCsi()
	ResetDownwardApi()
	ResetEmptyDir()
	ResetEphemeral()
	ResetFc()
	ResetFlexVolume()
	ResetFlocker()
	ResetGcePersistentDisk()
	ResetGitRepo()
	ResetGlusterfs()
	ResetHostPath()
	ResetIscsi()
	ResetLocal()
	ResetName()
	ResetNfs()
	ResetPersistentVolumeClaim()
	ResetPhotonPersistentDisk()
	ResetProjected()
	ResetQuobyte()
	ResetRbd()
	ResetSecret()
	ResetVsphereVolume()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DaemonSetV1SpecTemplateSpecVolumeOutputReference
type jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() DaemonSetV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *DaemonSetV1SpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *DaemonSetV1SpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) AzureDisk() DaemonSetV1SpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) AzureDiskInput() *DaemonSetV1SpecTemplateSpecVolumeAzureDisk {
	var returns *DaemonSetV1SpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) AzureFile() DaemonSetV1SpecTemplateSpecVolumeAzureFileOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) AzureFileInput() *DaemonSetV1SpecTemplateSpecVolumeAzureFile {
	var returns *DaemonSetV1SpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) CephFs() DaemonSetV1SpecTemplateSpecVolumeCephFsOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) CephFsInput() *DaemonSetV1SpecTemplateSpecVolumeCephFs {
	var returns *DaemonSetV1SpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Cinder() DaemonSetV1SpecTemplateSpecVolumeCinderOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) CinderInput() *DaemonSetV1SpecTemplateSpecVolumeCinder {
	var returns *DaemonSetV1SpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ConfigMap() DaemonSetV1SpecTemplateSpecVolumeConfigMapOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ConfigMapInput() *DaemonSetV1SpecTemplateSpecVolumeConfigMap {
	var returns *DaemonSetV1SpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Csi() DaemonSetV1SpecTemplateSpecVolumeCsiOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) CsiInput() *DaemonSetV1SpecTemplateSpecVolumeCsi {
	var returns *DaemonSetV1SpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) DownwardApi() DaemonSetV1SpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) DownwardApiInput() *DaemonSetV1SpecTemplateSpecVolumeDownwardApi {
	var returns *DaemonSetV1SpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) EmptyDir() DaemonSetV1SpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) EmptyDirInput() *DaemonSetV1SpecTemplateSpecVolumeEmptyDir {
	var returns *DaemonSetV1SpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Ephemeral() DaemonSetV1SpecTemplateSpecVolumeEphemeralOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeEphemeralOutputReference
	_jsii_.Get(
		j,
		"ephemeral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) EphemeralInput() *DaemonSetV1SpecTemplateSpecVolumeEphemeral {
	var returns *DaemonSetV1SpecTemplateSpecVolumeEphemeral
	_jsii_.Get(
		j,
		"ephemeralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Fc() DaemonSetV1SpecTemplateSpecVolumeFcOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) FcInput() *DaemonSetV1SpecTemplateSpecVolumeFc {
	var returns *DaemonSetV1SpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) FlexVolume() DaemonSetV1SpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *DaemonSetV1SpecTemplateSpecVolumeFlexVolume {
	var returns *DaemonSetV1SpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Flocker() DaemonSetV1SpecTemplateSpecVolumeFlockerOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) FlockerInput() *DaemonSetV1SpecTemplateSpecVolumeFlocker {
	var returns *DaemonSetV1SpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GcePersistentDisk() DaemonSetV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *DaemonSetV1SpecTemplateSpecVolumeGcePersistentDisk {
	var returns *DaemonSetV1SpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GitRepo() DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GitRepoInput() *DaemonSetV1SpecTemplateSpecVolumeGitRepo {
	var returns *DaemonSetV1SpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Glusterfs() DaemonSetV1SpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GlusterfsInput() *DaemonSetV1SpecTemplateSpecVolumeGlusterfs {
	var returns *DaemonSetV1SpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) HostPath() DaemonSetV1SpecTemplateSpecVolumeHostPathOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) HostPathInput() *DaemonSetV1SpecTemplateSpecVolumeHostPath {
	var returns *DaemonSetV1SpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Iscsi() DaemonSetV1SpecTemplateSpecVolumeIscsiOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) IscsiInput() *DaemonSetV1SpecTemplateSpecVolumeIscsi {
	var returns *DaemonSetV1SpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Local() DaemonSetV1SpecTemplateSpecVolumeLocalOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) LocalInput() *DaemonSetV1SpecTemplateSpecVolumeLocal {
	var returns *DaemonSetV1SpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Nfs() DaemonSetV1SpecTemplateSpecVolumeNfsOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) NfsInput() *DaemonSetV1SpecTemplateSpecVolumeNfs {
	var returns *DaemonSetV1SpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() DaemonSetV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *DaemonSetV1SpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *DaemonSetV1SpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Projected() DaemonSetV1SpecTemplateSpecVolumeProjectedList {
	var returns DaemonSetV1SpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Quobyte() DaemonSetV1SpecTemplateSpecVolumeQuobyteOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) QuobyteInput() *DaemonSetV1SpecTemplateSpecVolumeQuobyte {
	var returns *DaemonSetV1SpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Rbd() DaemonSetV1SpecTemplateSpecVolumeRbdOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) RbdInput() *DaemonSetV1SpecTemplateSpecVolumeRbd {
	var returns *DaemonSetV1SpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Secret() DaemonSetV1SpecTemplateSpecVolumeSecretOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) SecretInput() *DaemonSetV1SpecTemplateSpecVolumeSecret {
	var returns *DaemonSetV1SpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) VsphereVolume() DaemonSetV1SpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns DaemonSetV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *DaemonSetV1SpecTemplateSpecVolumeVsphereVolume {
	var returns *DaemonSetV1SpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewDaemonSetV1SpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DaemonSetV1SpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewDaemonSetV1SpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonSetV1.DaemonSetV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDaemonSetV1SpecTemplateSpecVolumeOutputReference_Override(d DaemonSetV1SpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.daemonSetV1.DaemonSetV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *DaemonSetV1SpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := d.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *DaemonSetV1SpecTemplateSpecVolumeAzureDisk) {
	if err := d.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutAzureFile(value *DaemonSetV1SpecTemplateSpecVolumeAzureFile) {
	if err := d.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutCephFs(value *DaemonSetV1SpecTemplateSpecVolumeCephFs) {
	if err := d.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCephFs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutCinder(value *DaemonSetV1SpecTemplateSpecVolumeCinder) {
	if err := d.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCinder",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutConfigMap(value *DaemonSetV1SpecTemplateSpecVolumeConfigMap) {
	if err := d.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutCsi(value *DaemonSetV1SpecTemplateSpecVolumeCsi) {
	if err := d.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *DaemonSetV1SpecTemplateSpecVolumeDownwardApi) {
	if err := d.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *DaemonSetV1SpecTemplateSpecVolumeEmptyDir) {
	if err := d.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutEphemeral(value *DaemonSetV1SpecTemplateSpecVolumeEphemeral) {
	if err := d.validatePutEphemeralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEphemeral",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutFc(value *DaemonSetV1SpecTemplateSpecVolumeFc) {
	if err := d.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFc",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *DaemonSetV1SpecTemplateSpecVolumeFlexVolume) {
	if err := d.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutFlocker(value *DaemonSetV1SpecTemplateSpecVolumeFlocker) {
	if err := d.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlocker",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *DaemonSetV1SpecTemplateSpecVolumeGcePersistentDisk) {
	if err := d.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutGitRepo(value *DaemonSetV1SpecTemplateSpecVolumeGitRepo) {
	if err := d.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *DaemonSetV1SpecTemplateSpecVolumeGlusterfs) {
	if err := d.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutHostPath(value *DaemonSetV1SpecTemplateSpecVolumeHostPath) {
	if err := d.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostPath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutIscsi(value *DaemonSetV1SpecTemplateSpecVolumeIscsi) {
	if err := d.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIscsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutLocal(value *DaemonSetV1SpecTemplateSpecVolumeLocal) {
	if err := d.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutNfs(value *DaemonSetV1SpecTemplateSpecVolumeNfs) {
	if err := d.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *DaemonSetV1SpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := d.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *DaemonSetV1SpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := d.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := d.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProjected",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutQuobyte(value *DaemonSetV1SpecTemplateSpecVolumeQuobyte) {
	if err := d.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutRbd(value *DaemonSetV1SpecTemplateSpecVolumeRbd) {
	if err := d.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRbd",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutSecret(value *DaemonSetV1SpecTemplateSpecVolumeSecret) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *DaemonSetV1SpecTemplateSpecVolumeVsphereVolume) {
	if err := d.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		d,
		"resetCephFs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		d,
		"resetCinder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		d,
		"resetCsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		d,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		d,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetEphemeral() {
	_jsii_.InvokeVoid(
		d,
		"resetEphemeral",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		d,
		"resetFc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		d,
		"resetFlocker",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		d,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		d,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		d,
		"resetHostPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		d,
		"resetLocal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		d,
		"resetNfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		d,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		d,
		"resetProjected",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		d,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		d,
		"resetRbd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DaemonSetV1SpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

