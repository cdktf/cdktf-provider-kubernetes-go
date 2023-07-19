package deploymentv1

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v8/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v8/deploymentv1/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentV1SpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() DeploymentV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *DeploymentV1SpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() DeploymentV1SpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *DeploymentV1SpecTemplateSpecVolumeAzureDisk
	AzureFile() DeploymentV1SpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *DeploymentV1SpecTemplateSpecVolumeAzureFile
	CephFs() DeploymentV1SpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *DeploymentV1SpecTemplateSpecVolumeCephFs
	Cinder() DeploymentV1SpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *DeploymentV1SpecTemplateSpecVolumeCinder
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
	ConfigMap() DeploymentV1SpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *DeploymentV1SpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() DeploymentV1SpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *DeploymentV1SpecTemplateSpecVolumeCsi
	DownwardApi() DeploymentV1SpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *DeploymentV1SpecTemplateSpecVolumeDownwardApi
	EmptyDir() DeploymentV1SpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *DeploymentV1SpecTemplateSpecVolumeEmptyDir
	Fc() DeploymentV1SpecTemplateSpecVolumeFcOutputReference
	FcInput() *DeploymentV1SpecTemplateSpecVolumeFc
	FlexVolume() DeploymentV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *DeploymentV1SpecTemplateSpecVolumeFlexVolume
	Flocker() DeploymentV1SpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *DeploymentV1SpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() DeploymentV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *DeploymentV1SpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() DeploymentV1SpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *DeploymentV1SpecTemplateSpecVolumeGitRepo
	Glusterfs() DeploymentV1SpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *DeploymentV1SpecTemplateSpecVolumeGlusterfs
	HostPath() DeploymentV1SpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *DeploymentV1SpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() DeploymentV1SpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *DeploymentV1SpecTemplateSpecVolumeIscsi
	Local() DeploymentV1SpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *DeploymentV1SpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() DeploymentV1SpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *DeploymentV1SpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() DeploymentV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *DeploymentV1SpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() DeploymentV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *DeploymentV1SpecTemplateSpecVolumePhotonPersistentDisk
	Projected() DeploymentV1SpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() DeploymentV1SpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *DeploymentV1SpecTemplateSpecVolumeQuobyte
	Rbd() DeploymentV1SpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *DeploymentV1SpecTemplateSpecVolumeRbd
	Secret() DeploymentV1SpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *DeploymentV1SpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() DeploymentV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *DeploymentV1SpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *DeploymentV1SpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *DeploymentV1SpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *DeploymentV1SpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *DeploymentV1SpecTemplateSpecVolumeCephFs)
	PutCinder(value *DeploymentV1SpecTemplateSpecVolumeCinder)
	PutConfigMap(value *DeploymentV1SpecTemplateSpecVolumeConfigMap)
	PutCsi(value *DeploymentV1SpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *DeploymentV1SpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *DeploymentV1SpecTemplateSpecVolumeEmptyDir)
	PutFc(value *DeploymentV1SpecTemplateSpecVolumeFc)
	PutFlexVolume(value *DeploymentV1SpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *DeploymentV1SpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *DeploymentV1SpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *DeploymentV1SpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *DeploymentV1SpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *DeploymentV1SpecTemplateSpecVolumeHostPath)
	PutIscsi(value *DeploymentV1SpecTemplateSpecVolumeIscsi)
	PutLocal(value *DeploymentV1SpecTemplateSpecVolumeLocal)
	PutNfs(value *DeploymentV1SpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *DeploymentV1SpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *DeploymentV1SpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *DeploymentV1SpecTemplateSpecVolumeQuobyte)
	PutRbd(value *DeploymentV1SpecTemplateSpecVolumeRbd)
	PutSecret(value *DeploymentV1SpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *DeploymentV1SpecTemplateSpecVolumeVsphereVolume)
	ResetAwsElasticBlockStore()
	ResetAzureDisk()
	ResetAzureFile()
	ResetCephFs()
	ResetCinder()
	ResetConfigMap()
	ResetCsi()
	ResetDownwardApi()
	ResetEmptyDir()
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

// The jsii proxy struct for DeploymentV1SpecTemplateSpecVolumeOutputReference
type jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() DeploymentV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *DeploymentV1SpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *DeploymentV1SpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) AzureDisk() DeploymentV1SpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) AzureDiskInput() *DeploymentV1SpecTemplateSpecVolumeAzureDisk {
	var returns *DeploymentV1SpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) AzureFile() DeploymentV1SpecTemplateSpecVolumeAzureFileOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) AzureFileInput() *DeploymentV1SpecTemplateSpecVolumeAzureFile {
	var returns *DeploymentV1SpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) CephFs() DeploymentV1SpecTemplateSpecVolumeCephFsOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) CephFsInput() *DeploymentV1SpecTemplateSpecVolumeCephFs {
	var returns *DeploymentV1SpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Cinder() DeploymentV1SpecTemplateSpecVolumeCinderOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) CinderInput() *DeploymentV1SpecTemplateSpecVolumeCinder {
	var returns *DeploymentV1SpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ConfigMap() DeploymentV1SpecTemplateSpecVolumeConfigMapOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ConfigMapInput() *DeploymentV1SpecTemplateSpecVolumeConfigMap {
	var returns *DeploymentV1SpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Csi() DeploymentV1SpecTemplateSpecVolumeCsiOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) CsiInput() *DeploymentV1SpecTemplateSpecVolumeCsi {
	var returns *DeploymentV1SpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) DownwardApi() DeploymentV1SpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) DownwardApiInput() *DeploymentV1SpecTemplateSpecVolumeDownwardApi {
	var returns *DeploymentV1SpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) EmptyDir() DeploymentV1SpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) EmptyDirInput() *DeploymentV1SpecTemplateSpecVolumeEmptyDir {
	var returns *DeploymentV1SpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Fc() DeploymentV1SpecTemplateSpecVolumeFcOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) FcInput() *DeploymentV1SpecTemplateSpecVolumeFc {
	var returns *DeploymentV1SpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) FlexVolume() DeploymentV1SpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *DeploymentV1SpecTemplateSpecVolumeFlexVolume {
	var returns *DeploymentV1SpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Flocker() DeploymentV1SpecTemplateSpecVolumeFlockerOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) FlockerInput() *DeploymentV1SpecTemplateSpecVolumeFlocker {
	var returns *DeploymentV1SpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GcePersistentDisk() DeploymentV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *DeploymentV1SpecTemplateSpecVolumeGcePersistentDisk {
	var returns *DeploymentV1SpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GitRepo() DeploymentV1SpecTemplateSpecVolumeGitRepoOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GitRepoInput() *DeploymentV1SpecTemplateSpecVolumeGitRepo {
	var returns *DeploymentV1SpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Glusterfs() DeploymentV1SpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GlusterfsInput() *DeploymentV1SpecTemplateSpecVolumeGlusterfs {
	var returns *DeploymentV1SpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) HostPath() DeploymentV1SpecTemplateSpecVolumeHostPathOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) HostPathInput() *DeploymentV1SpecTemplateSpecVolumeHostPath {
	var returns *DeploymentV1SpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Iscsi() DeploymentV1SpecTemplateSpecVolumeIscsiOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) IscsiInput() *DeploymentV1SpecTemplateSpecVolumeIscsi {
	var returns *DeploymentV1SpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Local() DeploymentV1SpecTemplateSpecVolumeLocalOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) LocalInput() *DeploymentV1SpecTemplateSpecVolumeLocal {
	var returns *DeploymentV1SpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Nfs() DeploymentV1SpecTemplateSpecVolumeNfsOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) NfsInput() *DeploymentV1SpecTemplateSpecVolumeNfs {
	var returns *DeploymentV1SpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() DeploymentV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *DeploymentV1SpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *DeploymentV1SpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() DeploymentV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *DeploymentV1SpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *DeploymentV1SpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Projected() DeploymentV1SpecTemplateSpecVolumeProjectedList {
	var returns DeploymentV1SpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Quobyte() DeploymentV1SpecTemplateSpecVolumeQuobyteOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) QuobyteInput() *DeploymentV1SpecTemplateSpecVolumeQuobyte {
	var returns *DeploymentV1SpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Rbd() DeploymentV1SpecTemplateSpecVolumeRbdOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) RbdInput() *DeploymentV1SpecTemplateSpecVolumeRbd {
	var returns *DeploymentV1SpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Secret() DeploymentV1SpecTemplateSpecVolumeSecretOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) SecretInput() *DeploymentV1SpecTemplateSpecVolumeSecret {
	var returns *DeploymentV1SpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) VsphereVolume() DeploymentV1SpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns DeploymentV1SpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *DeploymentV1SpecTemplateSpecVolumeVsphereVolume {
	var returns *DeploymentV1SpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewDeploymentV1SpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DeploymentV1SpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentV1SpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDeploymentV1SpecTemplateSpecVolumeOutputReference_Override(d DeploymentV1SpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deploymentV1.DeploymentV1SpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *DeploymentV1SpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := d.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *DeploymentV1SpecTemplateSpecVolumeAzureDisk) {
	if err := d.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutAzureFile(value *DeploymentV1SpecTemplateSpecVolumeAzureFile) {
	if err := d.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutCephFs(value *DeploymentV1SpecTemplateSpecVolumeCephFs) {
	if err := d.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCephFs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutCinder(value *DeploymentV1SpecTemplateSpecVolumeCinder) {
	if err := d.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCinder",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutConfigMap(value *DeploymentV1SpecTemplateSpecVolumeConfigMap) {
	if err := d.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutCsi(value *DeploymentV1SpecTemplateSpecVolumeCsi) {
	if err := d.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *DeploymentV1SpecTemplateSpecVolumeDownwardApi) {
	if err := d.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *DeploymentV1SpecTemplateSpecVolumeEmptyDir) {
	if err := d.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutFc(value *DeploymentV1SpecTemplateSpecVolumeFc) {
	if err := d.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFc",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *DeploymentV1SpecTemplateSpecVolumeFlexVolume) {
	if err := d.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutFlocker(value *DeploymentV1SpecTemplateSpecVolumeFlocker) {
	if err := d.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlocker",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *DeploymentV1SpecTemplateSpecVolumeGcePersistentDisk) {
	if err := d.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutGitRepo(value *DeploymentV1SpecTemplateSpecVolumeGitRepo) {
	if err := d.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *DeploymentV1SpecTemplateSpecVolumeGlusterfs) {
	if err := d.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutHostPath(value *DeploymentV1SpecTemplateSpecVolumeHostPath) {
	if err := d.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostPath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutIscsi(value *DeploymentV1SpecTemplateSpecVolumeIscsi) {
	if err := d.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIscsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutLocal(value *DeploymentV1SpecTemplateSpecVolumeLocal) {
	if err := d.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutNfs(value *DeploymentV1SpecTemplateSpecVolumeNfs) {
	if err := d.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *DeploymentV1SpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := d.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *DeploymentV1SpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := d.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := d.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProjected",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutQuobyte(value *DeploymentV1SpecTemplateSpecVolumeQuobyte) {
	if err := d.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutRbd(value *DeploymentV1SpecTemplateSpecVolumeRbd) {
	if err := d.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRbd",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutSecret(value *DeploymentV1SpecTemplateSpecVolumeSecret) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *DeploymentV1SpecTemplateSpecVolumeVsphereVolume) {
	if err := d.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		d,
		"resetCephFs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		d,
		"resetCinder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		d,
		"resetCsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		d,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		d,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		d,
		"resetFc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		d,
		"resetFlocker",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		d,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		d,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		d,
		"resetHostPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		d,
		"resetLocal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		d,
		"resetNfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		d,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		d,
		"resetProjected",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		d,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		d,
		"resetRbd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentV1SpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

