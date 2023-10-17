// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulsetv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v10/statefulsetv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulSetV1SpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() StatefulSetV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *StatefulSetV1SpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() StatefulSetV1SpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *StatefulSetV1SpecTemplateSpecVolumeAzureDisk
	AzureFile() StatefulSetV1SpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *StatefulSetV1SpecTemplateSpecVolumeAzureFile
	CephFs() StatefulSetV1SpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *StatefulSetV1SpecTemplateSpecVolumeCephFs
	Cinder() StatefulSetV1SpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *StatefulSetV1SpecTemplateSpecVolumeCinder
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
	ConfigMap() StatefulSetV1SpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *StatefulSetV1SpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() StatefulSetV1SpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *StatefulSetV1SpecTemplateSpecVolumeCsi
	DownwardApi() StatefulSetV1SpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *StatefulSetV1SpecTemplateSpecVolumeDownwardApi
	EmptyDir() StatefulSetV1SpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *StatefulSetV1SpecTemplateSpecVolumeEmptyDir
	Ephemeral() StatefulSetV1SpecTemplateSpecVolumeEphemeralOutputReference
	EphemeralInput() *StatefulSetV1SpecTemplateSpecVolumeEphemeral
	Fc() StatefulSetV1SpecTemplateSpecVolumeFcOutputReference
	FcInput() *StatefulSetV1SpecTemplateSpecVolumeFc
	FlexVolume() StatefulSetV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *StatefulSetV1SpecTemplateSpecVolumeFlexVolume
	Flocker() StatefulSetV1SpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *StatefulSetV1SpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() StatefulSetV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *StatefulSetV1SpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() StatefulSetV1SpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *StatefulSetV1SpecTemplateSpecVolumeGitRepo
	Glusterfs() StatefulSetV1SpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *StatefulSetV1SpecTemplateSpecVolumeGlusterfs
	HostPath() StatefulSetV1SpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *StatefulSetV1SpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() StatefulSetV1SpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *StatefulSetV1SpecTemplateSpecVolumeIscsi
	Local() StatefulSetV1SpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *StatefulSetV1SpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() StatefulSetV1SpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *StatefulSetV1SpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() StatefulSetV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *StatefulSetV1SpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() StatefulSetV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *StatefulSetV1SpecTemplateSpecVolumePhotonPersistentDisk
	Projected() StatefulSetV1SpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() StatefulSetV1SpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *StatefulSetV1SpecTemplateSpecVolumeQuobyte
	Rbd() StatefulSetV1SpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *StatefulSetV1SpecTemplateSpecVolumeRbd
	Secret() StatefulSetV1SpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *StatefulSetV1SpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() StatefulSetV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *StatefulSetV1SpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *StatefulSetV1SpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *StatefulSetV1SpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *StatefulSetV1SpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *StatefulSetV1SpecTemplateSpecVolumeCephFs)
	PutCinder(value *StatefulSetV1SpecTemplateSpecVolumeCinder)
	PutConfigMap(value *StatefulSetV1SpecTemplateSpecVolumeConfigMap)
	PutCsi(value *StatefulSetV1SpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *StatefulSetV1SpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *StatefulSetV1SpecTemplateSpecVolumeEmptyDir)
	PutEphemeral(value *StatefulSetV1SpecTemplateSpecVolumeEphemeral)
	PutFc(value *StatefulSetV1SpecTemplateSpecVolumeFc)
	PutFlexVolume(value *StatefulSetV1SpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *StatefulSetV1SpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *StatefulSetV1SpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *StatefulSetV1SpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *StatefulSetV1SpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *StatefulSetV1SpecTemplateSpecVolumeHostPath)
	PutIscsi(value *StatefulSetV1SpecTemplateSpecVolumeIscsi)
	PutLocal(value *StatefulSetV1SpecTemplateSpecVolumeLocal)
	PutNfs(value *StatefulSetV1SpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *StatefulSetV1SpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *StatefulSetV1SpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *StatefulSetV1SpecTemplateSpecVolumeQuobyte)
	PutRbd(value *StatefulSetV1SpecTemplateSpecVolumeRbd)
	PutSecret(value *StatefulSetV1SpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *StatefulSetV1SpecTemplateSpecVolumeVsphereVolume)
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
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StatefulSetV1SpecTemplateSpecVolumeOutputReference
type jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() StatefulSetV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *StatefulSetV1SpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *StatefulSetV1SpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) AzureDisk() StatefulSetV1SpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) AzureDiskInput() *StatefulSetV1SpecTemplateSpecVolumeAzureDisk {
	var returns *StatefulSetV1SpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) AzureFile() StatefulSetV1SpecTemplateSpecVolumeAzureFileOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) AzureFileInput() *StatefulSetV1SpecTemplateSpecVolumeAzureFile {
	var returns *StatefulSetV1SpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) CephFs() StatefulSetV1SpecTemplateSpecVolumeCephFsOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) CephFsInput() *StatefulSetV1SpecTemplateSpecVolumeCephFs {
	var returns *StatefulSetV1SpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Cinder() StatefulSetV1SpecTemplateSpecVolumeCinderOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) CinderInput() *StatefulSetV1SpecTemplateSpecVolumeCinder {
	var returns *StatefulSetV1SpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ConfigMap() StatefulSetV1SpecTemplateSpecVolumeConfigMapOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ConfigMapInput() *StatefulSetV1SpecTemplateSpecVolumeConfigMap {
	var returns *StatefulSetV1SpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Csi() StatefulSetV1SpecTemplateSpecVolumeCsiOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) CsiInput() *StatefulSetV1SpecTemplateSpecVolumeCsi {
	var returns *StatefulSetV1SpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) DownwardApi() StatefulSetV1SpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) DownwardApiInput() *StatefulSetV1SpecTemplateSpecVolumeDownwardApi {
	var returns *StatefulSetV1SpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) EmptyDir() StatefulSetV1SpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) EmptyDirInput() *StatefulSetV1SpecTemplateSpecVolumeEmptyDir {
	var returns *StatefulSetV1SpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Ephemeral() StatefulSetV1SpecTemplateSpecVolumeEphemeralOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeEphemeralOutputReference
	_jsii_.Get(
		j,
		"ephemeral",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) EphemeralInput() *StatefulSetV1SpecTemplateSpecVolumeEphemeral {
	var returns *StatefulSetV1SpecTemplateSpecVolumeEphemeral
	_jsii_.Get(
		j,
		"ephemeralInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Fc() StatefulSetV1SpecTemplateSpecVolumeFcOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) FcInput() *StatefulSetV1SpecTemplateSpecVolumeFc {
	var returns *StatefulSetV1SpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) FlexVolume() StatefulSetV1SpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *StatefulSetV1SpecTemplateSpecVolumeFlexVolume {
	var returns *StatefulSetV1SpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Flocker() StatefulSetV1SpecTemplateSpecVolumeFlockerOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) FlockerInput() *StatefulSetV1SpecTemplateSpecVolumeFlocker {
	var returns *StatefulSetV1SpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GcePersistentDisk() StatefulSetV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *StatefulSetV1SpecTemplateSpecVolumeGcePersistentDisk {
	var returns *StatefulSetV1SpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GitRepo() StatefulSetV1SpecTemplateSpecVolumeGitRepoOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GitRepoInput() *StatefulSetV1SpecTemplateSpecVolumeGitRepo {
	var returns *StatefulSetV1SpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Glusterfs() StatefulSetV1SpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GlusterfsInput() *StatefulSetV1SpecTemplateSpecVolumeGlusterfs {
	var returns *StatefulSetV1SpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) HostPath() StatefulSetV1SpecTemplateSpecVolumeHostPathOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) HostPathInput() *StatefulSetV1SpecTemplateSpecVolumeHostPath {
	var returns *StatefulSetV1SpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Iscsi() StatefulSetV1SpecTemplateSpecVolumeIscsiOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) IscsiInput() *StatefulSetV1SpecTemplateSpecVolumeIscsi {
	var returns *StatefulSetV1SpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Local() StatefulSetV1SpecTemplateSpecVolumeLocalOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) LocalInput() *StatefulSetV1SpecTemplateSpecVolumeLocal {
	var returns *StatefulSetV1SpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Nfs() StatefulSetV1SpecTemplateSpecVolumeNfsOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) NfsInput() *StatefulSetV1SpecTemplateSpecVolumeNfs {
	var returns *StatefulSetV1SpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() StatefulSetV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *StatefulSetV1SpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *StatefulSetV1SpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() StatefulSetV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *StatefulSetV1SpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *StatefulSetV1SpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Projected() StatefulSetV1SpecTemplateSpecVolumeProjectedList {
	var returns StatefulSetV1SpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Quobyte() StatefulSetV1SpecTemplateSpecVolumeQuobyteOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) QuobyteInput() *StatefulSetV1SpecTemplateSpecVolumeQuobyte {
	var returns *StatefulSetV1SpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Rbd() StatefulSetV1SpecTemplateSpecVolumeRbdOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) RbdInput() *StatefulSetV1SpecTemplateSpecVolumeRbd {
	var returns *StatefulSetV1SpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Secret() StatefulSetV1SpecTemplateSpecVolumeSecretOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) SecretInput() *StatefulSetV1SpecTemplateSpecVolumeSecret {
	var returns *StatefulSetV1SpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) VsphereVolume() StatefulSetV1SpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns StatefulSetV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *StatefulSetV1SpecTemplateSpecVolumeVsphereVolume {
	var returns *StatefulSetV1SpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewStatefulSetV1SpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StatefulSetV1SpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewStatefulSetV1SpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSetV1.StatefulSetV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStatefulSetV1SpecTemplateSpecVolumeOutputReference_Override(s StatefulSetV1SpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.statefulSetV1.StatefulSetV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *StatefulSetV1SpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := s.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *StatefulSetV1SpecTemplateSpecVolumeAzureDisk) {
	if err := s.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutAzureFile(value *StatefulSetV1SpecTemplateSpecVolumeAzureFile) {
	if err := s.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutCephFs(value *StatefulSetV1SpecTemplateSpecVolumeCephFs) {
	if err := s.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCephFs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutCinder(value *StatefulSetV1SpecTemplateSpecVolumeCinder) {
	if err := s.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCinder",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutConfigMap(value *StatefulSetV1SpecTemplateSpecVolumeConfigMap) {
	if err := s.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutCsi(value *StatefulSetV1SpecTemplateSpecVolumeCsi) {
	if err := s.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCsi",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *StatefulSetV1SpecTemplateSpecVolumeDownwardApi) {
	if err := s.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *StatefulSetV1SpecTemplateSpecVolumeEmptyDir) {
	if err := s.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutEphemeral(value *StatefulSetV1SpecTemplateSpecVolumeEphemeral) {
	if err := s.validatePutEphemeralParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEphemeral",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutFc(value *StatefulSetV1SpecTemplateSpecVolumeFc) {
	if err := s.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFc",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *StatefulSetV1SpecTemplateSpecVolumeFlexVolume) {
	if err := s.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutFlocker(value *StatefulSetV1SpecTemplateSpecVolumeFlocker) {
	if err := s.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFlocker",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *StatefulSetV1SpecTemplateSpecVolumeGcePersistentDisk) {
	if err := s.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutGitRepo(value *StatefulSetV1SpecTemplateSpecVolumeGitRepo) {
	if err := s.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *StatefulSetV1SpecTemplateSpecVolumeGlusterfs) {
	if err := s.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutHostPath(value *StatefulSetV1SpecTemplateSpecVolumeHostPath) {
	if err := s.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHostPath",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutIscsi(value *StatefulSetV1SpecTemplateSpecVolumeIscsi) {
	if err := s.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIscsi",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutLocal(value *StatefulSetV1SpecTemplateSpecVolumeLocal) {
	if err := s.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLocal",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutNfs(value *StatefulSetV1SpecTemplateSpecVolumeNfs) {
	if err := s.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNfs",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *StatefulSetV1SpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := s.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *StatefulSetV1SpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := s.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := s.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putProjected",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutQuobyte(value *StatefulSetV1SpecTemplateSpecVolumeQuobyte) {
	if err := s.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutRbd(value *StatefulSetV1SpecTemplateSpecVolumeRbd) {
	if err := s.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRbd",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutSecret(value *StatefulSetV1SpecTemplateSpecVolumeSecret) {
	if err := s.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSecret",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *StatefulSetV1SpecTemplateSpecVolumeVsphereVolume) {
	if err := s.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		s,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		s,
		"resetCephFs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		s,
		"resetCinder",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		s,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		s,
		"resetCsi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		s,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		s,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetEphemeral() {
	_jsii_.InvokeVoid(
		s,
		"resetEphemeral",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		s,
		"resetFc",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		s,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		s,
		"resetFlocker",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		s,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		s,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		s,
		"resetHostPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		s,
		"resetIscsi",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		s,
		"resetLocal",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		s,
		"resetNfs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		s,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		s,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		s,
		"resetProjected",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		s,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		s,
		"resetRbd",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		s,
		"resetSecret",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		s,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StatefulSetV1SpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

