package persistentvolumev1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v4/persistentvolumev1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PersistentVolumeV1SpecPersistentVolumeSourceOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore
	AzureDisk() PersistentVolumeV1SpecPersistentVolumeSourceAzureDiskOutputReference
	AzureDiskInput() *PersistentVolumeV1SpecPersistentVolumeSourceAzureDisk
	AzureFile() PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference
	AzureFileInput() *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile
	CephFs() PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference
	CephFsInput() *PersistentVolumeV1SpecPersistentVolumeSourceCephFs
	Cinder() PersistentVolumeV1SpecPersistentVolumeSourceCinderOutputReference
	CinderInput() *PersistentVolumeV1SpecPersistentVolumeSourceCinder
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
	Csi() PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference
	CsiInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsi
	Fc() PersistentVolumeV1SpecPersistentVolumeSourceFcOutputReference
	FcInput() *PersistentVolumeV1SpecPersistentVolumeSourceFc
	FlexVolume() PersistentVolumeV1SpecPersistentVolumeSourceFlexVolumeOutputReference
	FlexVolumeInput() *PersistentVolumeV1SpecPersistentVolumeSourceFlexVolume
	Flocker() PersistentVolumeV1SpecPersistentVolumeSourceFlockerOutputReference
	FlockerInput() *PersistentVolumeV1SpecPersistentVolumeSourceFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDiskOutputReference
	GcePersistentDiskInput() *PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk
	Glusterfs() PersistentVolumeV1SpecPersistentVolumeSourceGlusterfsOutputReference
	GlusterfsInput() *PersistentVolumeV1SpecPersistentVolumeSourceGlusterfs
	HostPath() PersistentVolumeV1SpecPersistentVolumeSourceHostPathOutputReference
	HostPathInput() *PersistentVolumeV1SpecPersistentVolumeSourceHostPath
	InternalValue() *PersistentVolumeV1SpecPersistentVolumeSource
	SetInternalValue(val *PersistentVolumeV1SpecPersistentVolumeSource)
	Iscsi() PersistentVolumeV1SpecPersistentVolumeSourceIscsiOutputReference
	IscsiInput() *PersistentVolumeV1SpecPersistentVolumeSourceIscsi
	Local() PersistentVolumeV1SpecPersistentVolumeSourceLocalOutputReference
	LocalInput() *PersistentVolumeV1SpecPersistentVolumeSourceLocal
	Nfs() PersistentVolumeV1SpecPersistentVolumeSourceNfsOutputReference
	NfsInput() *PersistentVolumeV1SpecPersistentVolumeSourceNfs
	PhotonPersistentDisk() PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk
	Quobyte() PersistentVolumeV1SpecPersistentVolumeSourceQuobyteOutputReference
	QuobyteInput() *PersistentVolumeV1SpecPersistentVolumeSourceQuobyte
	Rbd() PersistentVolumeV1SpecPersistentVolumeSourceRbdOutputReference
	RbdInput() *PersistentVolumeV1SpecPersistentVolumeSourceRbd
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolumeOutputReference
	VsphereVolumeInput() *PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume
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
	PutAwsElasticBlockStore(value *PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore)
	PutAzureDisk(value *PersistentVolumeV1SpecPersistentVolumeSourceAzureDisk)
	PutAzureFile(value *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile)
	PutCephFs(value *PersistentVolumeV1SpecPersistentVolumeSourceCephFs)
	PutCinder(value *PersistentVolumeV1SpecPersistentVolumeSourceCinder)
	PutCsi(value *PersistentVolumeV1SpecPersistentVolumeSourceCsi)
	PutFc(value *PersistentVolumeV1SpecPersistentVolumeSourceFc)
	PutFlexVolume(value *PersistentVolumeV1SpecPersistentVolumeSourceFlexVolume)
	PutFlocker(value *PersistentVolumeV1SpecPersistentVolumeSourceFlocker)
	PutGcePersistentDisk(value *PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk)
	PutGlusterfs(value *PersistentVolumeV1SpecPersistentVolumeSourceGlusterfs)
	PutHostPath(value *PersistentVolumeV1SpecPersistentVolumeSourceHostPath)
	PutIscsi(value *PersistentVolumeV1SpecPersistentVolumeSourceIscsi)
	PutLocal(value *PersistentVolumeV1SpecPersistentVolumeSourceLocal)
	PutNfs(value *PersistentVolumeV1SpecPersistentVolumeSourceNfs)
	PutPhotonPersistentDisk(value *PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk)
	PutQuobyte(value *PersistentVolumeV1SpecPersistentVolumeSourceQuobyte)
	PutRbd(value *PersistentVolumeV1SpecPersistentVolumeSourceRbd)
	PutVsphereVolume(value *PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume)
	ResetAwsElasticBlockStore()
	ResetAzureDisk()
	ResetAzureFile()
	ResetCephFs()
	ResetCinder()
	ResetCsi()
	ResetFc()
	ResetFlexVolume()
	ResetFlocker()
	ResetGcePersistentDisk()
	ResetGlusterfs()
	ResetHostPath()
	ResetIscsi()
	ResetLocal()
	ResetNfs()
	ResetPhotonPersistentDisk()
	ResetQuobyte()
	ResetRbd()
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

// The jsii proxy struct for PersistentVolumeV1SpecPersistentVolumeSourceOutputReference
type jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AwsElasticBlockStore() PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AwsElasticBlockStoreInput() *PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AzureDisk() PersistentVolumeV1SpecPersistentVolumeSourceAzureDiskOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AzureDiskInput() *PersistentVolumeV1SpecPersistentVolumeSourceAzureDisk {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AzureFile() PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) AzureFileInput() *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CephFs() PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CephFsInput() *PersistentVolumeV1SpecPersistentVolumeSourceCephFs {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Cinder() PersistentVolumeV1SpecPersistentVolumeSourceCinderOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CinderInput() *PersistentVolumeV1SpecPersistentVolumeSourceCinder {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Csi() PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) CsiInput() *PersistentVolumeV1SpecPersistentVolumeSourceCsi {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Fc() PersistentVolumeV1SpecPersistentVolumeSourceFcOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) FcInput() *PersistentVolumeV1SpecPersistentVolumeSourceFc {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) FlexVolume() PersistentVolumeV1SpecPersistentVolumeSourceFlexVolumeOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) FlexVolumeInput() *PersistentVolumeV1SpecPersistentVolumeSourceFlexVolume {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Flocker() PersistentVolumeV1SpecPersistentVolumeSourceFlockerOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) FlockerInput() *PersistentVolumeV1SpecPersistentVolumeSourceFlocker {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GcePersistentDisk() PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDiskOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GcePersistentDiskInput() *PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Glusterfs() PersistentVolumeV1SpecPersistentVolumeSourceGlusterfsOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GlusterfsInput() *PersistentVolumeV1SpecPersistentVolumeSourceGlusterfs {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) HostPath() PersistentVolumeV1SpecPersistentVolumeSourceHostPathOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) HostPathInput() *PersistentVolumeV1SpecPersistentVolumeSourceHostPath {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) InternalValue() *PersistentVolumeV1SpecPersistentVolumeSource {
	var returns *PersistentVolumeV1SpecPersistentVolumeSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Iscsi() PersistentVolumeV1SpecPersistentVolumeSourceIscsiOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) IscsiInput() *PersistentVolumeV1SpecPersistentVolumeSourceIscsi {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Local() PersistentVolumeV1SpecPersistentVolumeSourceLocalOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) LocalInput() *PersistentVolumeV1SpecPersistentVolumeSourceLocal {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Nfs() PersistentVolumeV1SpecPersistentVolumeSourceNfsOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) NfsInput() *PersistentVolumeV1SpecPersistentVolumeSourceNfs {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PhotonPersistentDisk() PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDiskOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PhotonPersistentDiskInput() *PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Quobyte() PersistentVolumeV1SpecPersistentVolumeSourceQuobyteOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) QuobyteInput() *PersistentVolumeV1SpecPersistentVolumeSourceQuobyte {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Rbd() PersistentVolumeV1SpecPersistentVolumeSourceRbdOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) RbdInput() *PersistentVolumeV1SpecPersistentVolumeSourceRbd {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) VsphereVolume() PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolumeOutputReference {
	var returns PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) VsphereVolumeInput() *PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume {
	var returns *PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewPersistentVolumeV1SpecPersistentVolumeSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PersistentVolumeV1SpecPersistentVolumeSourceOutputReference {
	_init_.Initialize()

	if err := validateNewPersistentVolumeV1SpecPersistentVolumeSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecPersistentVolumeSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPersistentVolumeV1SpecPersistentVolumeSourceOutputReference_Override(p PersistentVolumeV1SpecPersistentVolumeSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.persistentVolumeV1.PersistentVolumeV1SpecPersistentVolumeSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetInternalValue(val *PersistentVolumeV1SpecPersistentVolumeSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutAwsElasticBlockStore(value *PersistentVolumeV1SpecPersistentVolumeSourceAwsElasticBlockStore) {
	if err := p.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutAzureDisk(value *PersistentVolumeV1SpecPersistentVolumeSourceAzureDisk) {
	if err := p.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutAzureFile(value *PersistentVolumeV1SpecPersistentVolumeSourceAzureFile) {
	if err := p.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutCephFs(value *PersistentVolumeV1SpecPersistentVolumeSourceCephFs) {
	if err := p.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCephFs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutCinder(value *PersistentVolumeV1SpecPersistentVolumeSourceCinder) {
	if err := p.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCinder",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutCsi(value *PersistentVolumeV1SpecPersistentVolumeSourceCsi) {
	if err := p.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCsi",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutFc(value *PersistentVolumeV1SpecPersistentVolumeSourceFc) {
	if err := p.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFc",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutFlexVolume(value *PersistentVolumeV1SpecPersistentVolumeSourceFlexVolume) {
	if err := p.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutFlocker(value *PersistentVolumeV1SpecPersistentVolumeSourceFlocker) {
	if err := p.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFlocker",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutGcePersistentDisk(value *PersistentVolumeV1SpecPersistentVolumeSourceGcePersistentDisk) {
	if err := p.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutGlusterfs(value *PersistentVolumeV1SpecPersistentVolumeSourceGlusterfs) {
	if err := p.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutHostPath(value *PersistentVolumeV1SpecPersistentVolumeSourceHostPath) {
	if err := p.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostPath",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutIscsi(value *PersistentVolumeV1SpecPersistentVolumeSourceIscsi) {
	if err := p.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIscsi",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutLocal(value *PersistentVolumeV1SpecPersistentVolumeSourceLocal) {
	if err := p.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLocal",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutNfs(value *PersistentVolumeV1SpecPersistentVolumeSourceNfs) {
	if err := p.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNfs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutPhotonPersistentDisk(value *PersistentVolumeV1SpecPersistentVolumeSourcePhotonPersistentDisk) {
	if err := p.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutQuobyte(value *PersistentVolumeV1SpecPersistentVolumeSourceQuobyte) {
	if err := p.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutRbd(value *PersistentVolumeV1SpecPersistentVolumeSourceRbd) {
	if err := p.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRbd",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) PutVsphereVolume(value *PersistentVolumeV1SpecPersistentVolumeSourceVsphereVolume) {
	if err := p.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		p,
		"resetCephFs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		p,
		"resetCinder",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		p,
		"resetCsi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		p,
		"resetFc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		p,
		"resetFlocker",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		p,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		p,
		"resetIscsi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		p,
		"resetLocal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		p,
		"resetNfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		p,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		p,
		"resetRbd",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PersistentVolumeV1SpecPersistentVolumeSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

