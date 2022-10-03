package deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v3/jsii"

	"github.com/hashicorp/cdktf-provider-kubernetes-go/kubernetes/v3/deployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DeploymentSpecTemplateSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() DeploymentSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *DeploymentSpecTemplateSpecVolumeAwsElasticBlockStore
	AzureDisk() DeploymentSpecTemplateSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *DeploymentSpecTemplateSpecVolumeAzureDisk
	AzureFile() DeploymentSpecTemplateSpecVolumeAzureFileOutputReference
	AzureFileInput() *DeploymentSpecTemplateSpecVolumeAzureFile
	CephFs() DeploymentSpecTemplateSpecVolumeCephFsOutputReference
	CephFsInput() *DeploymentSpecTemplateSpecVolumeCephFs
	Cinder() DeploymentSpecTemplateSpecVolumeCinderOutputReference
	CinderInput() *DeploymentSpecTemplateSpecVolumeCinder
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
	ConfigMap() DeploymentSpecTemplateSpecVolumeConfigMapOutputReference
	ConfigMapInput() *DeploymentSpecTemplateSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() DeploymentSpecTemplateSpecVolumeCsiOutputReference
	CsiInput() *DeploymentSpecTemplateSpecVolumeCsi
	DownwardApi() DeploymentSpecTemplateSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *DeploymentSpecTemplateSpecVolumeDownwardApi
	EmptyDir() DeploymentSpecTemplateSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *DeploymentSpecTemplateSpecVolumeEmptyDir
	Fc() DeploymentSpecTemplateSpecVolumeFcOutputReference
	FcInput() *DeploymentSpecTemplateSpecVolumeFc
	FlexVolume() DeploymentSpecTemplateSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *DeploymentSpecTemplateSpecVolumeFlexVolume
	Flocker() DeploymentSpecTemplateSpecVolumeFlockerOutputReference
	FlockerInput() *DeploymentSpecTemplateSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() DeploymentSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *DeploymentSpecTemplateSpecVolumeGcePersistentDisk
	GitRepo() DeploymentSpecTemplateSpecVolumeGitRepoOutputReference
	GitRepoInput() *DeploymentSpecTemplateSpecVolumeGitRepo
	Glusterfs() DeploymentSpecTemplateSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *DeploymentSpecTemplateSpecVolumeGlusterfs
	HostPath() DeploymentSpecTemplateSpecVolumeHostPathOutputReference
	HostPathInput() *DeploymentSpecTemplateSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() DeploymentSpecTemplateSpecVolumeIscsiOutputReference
	IscsiInput() *DeploymentSpecTemplateSpecVolumeIscsi
	Local() DeploymentSpecTemplateSpecVolumeLocalOutputReference
	LocalInput() *DeploymentSpecTemplateSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() DeploymentSpecTemplateSpecVolumeNfsOutputReference
	NfsInput() *DeploymentSpecTemplateSpecVolumeNfs
	PersistentVolumeClaim() DeploymentSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *DeploymentSpecTemplateSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() DeploymentSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *DeploymentSpecTemplateSpecVolumePhotonPersistentDisk
	Projected() DeploymentSpecTemplateSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() DeploymentSpecTemplateSpecVolumeQuobyteOutputReference
	QuobyteInput() *DeploymentSpecTemplateSpecVolumeQuobyte
	Rbd() DeploymentSpecTemplateSpecVolumeRbdOutputReference
	RbdInput() *DeploymentSpecTemplateSpecVolumeRbd
	Secret() DeploymentSpecTemplateSpecVolumeSecretOutputReference
	SecretInput() *DeploymentSpecTemplateSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() DeploymentSpecTemplateSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *DeploymentSpecTemplateSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *DeploymentSpecTemplateSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *DeploymentSpecTemplateSpecVolumeAzureDisk)
	PutAzureFile(value *DeploymentSpecTemplateSpecVolumeAzureFile)
	PutCephFs(value *DeploymentSpecTemplateSpecVolumeCephFs)
	PutCinder(value *DeploymentSpecTemplateSpecVolumeCinder)
	PutConfigMap(value *DeploymentSpecTemplateSpecVolumeConfigMap)
	PutCsi(value *DeploymentSpecTemplateSpecVolumeCsi)
	PutDownwardApi(value *DeploymentSpecTemplateSpecVolumeDownwardApi)
	PutEmptyDir(value *DeploymentSpecTemplateSpecVolumeEmptyDir)
	PutFc(value *DeploymentSpecTemplateSpecVolumeFc)
	PutFlexVolume(value *DeploymentSpecTemplateSpecVolumeFlexVolume)
	PutFlocker(value *DeploymentSpecTemplateSpecVolumeFlocker)
	PutGcePersistentDisk(value *DeploymentSpecTemplateSpecVolumeGcePersistentDisk)
	PutGitRepo(value *DeploymentSpecTemplateSpecVolumeGitRepo)
	PutGlusterfs(value *DeploymentSpecTemplateSpecVolumeGlusterfs)
	PutHostPath(value *DeploymentSpecTemplateSpecVolumeHostPath)
	PutIscsi(value *DeploymentSpecTemplateSpecVolumeIscsi)
	PutLocal(value *DeploymentSpecTemplateSpecVolumeLocal)
	PutNfs(value *DeploymentSpecTemplateSpecVolumeNfs)
	PutPersistentVolumeClaim(value *DeploymentSpecTemplateSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *DeploymentSpecTemplateSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *DeploymentSpecTemplateSpecVolumeQuobyte)
	PutRbd(value *DeploymentSpecTemplateSpecVolumeRbd)
	PutSecret(value *DeploymentSpecTemplateSpecVolumeSecret)
	PutVsphereVolume(value *DeploymentSpecTemplateSpecVolumeVsphereVolume)
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

// The jsii proxy struct for DeploymentSpecTemplateSpecVolumeOutputReference
type jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStore() DeploymentSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) AwsElasticBlockStoreInput() *DeploymentSpecTemplateSpecVolumeAwsElasticBlockStore {
	var returns *DeploymentSpecTemplateSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) AzureDisk() DeploymentSpecTemplateSpecVolumeAzureDiskOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) AzureDiskInput() *DeploymentSpecTemplateSpecVolumeAzureDisk {
	var returns *DeploymentSpecTemplateSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) AzureFile() DeploymentSpecTemplateSpecVolumeAzureFileOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) AzureFileInput() *DeploymentSpecTemplateSpecVolumeAzureFile {
	var returns *DeploymentSpecTemplateSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) CephFs() DeploymentSpecTemplateSpecVolumeCephFsOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) CephFsInput() *DeploymentSpecTemplateSpecVolumeCephFs {
	var returns *DeploymentSpecTemplateSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Cinder() DeploymentSpecTemplateSpecVolumeCinderOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) CinderInput() *DeploymentSpecTemplateSpecVolumeCinder {
	var returns *DeploymentSpecTemplateSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ConfigMap() DeploymentSpecTemplateSpecVolumeConfigMapOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ConfigMapInput() *DeploymentSpecTemplateSpecVolumeConfigMap {
	var returns *DeploymentSpecTemplateSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Csi() DeploymentSpecTemplateSpecVolumeCsiOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) CsiInput() *DeploymentSpecTemplateSpecVolumeCsi {
	var returns *DeploymentSpecTemplateSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) DownwardApi() DeploymentSpecTemplateSpecVolumeDownwardApiOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) DownwardApiInput() *DeploymentSpecTemplateSpecVolumeDownwardApi {
	var returns *DeploymentSpecTemplateSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) EmptyDir() DeploymentSpecTemplateSpecVolumeEmptyDirOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) EmptyDirInput() *DeploymentSpecTemplateSpecVolumeEmptyDir {
	var returns *DeploymentSpecTemplateSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Fc() DeploymentSpecTemplateSpecVolumeFcOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) FcInput() *DeploymentSpecTemplateSpecVolumeFc {
	var returns *DeploymentSpecTemplateSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) FlexVolume() DeploymentSpecTemplateSpecVolumeFlexVolumeOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) FlexVolumeInput() *DeploymentSpecTemplateSpecVolumeFlexVolume {
	var returns *DeploymentSpecTemplateSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Flocker() DeploymentSpecTemplateSpecVolumeFlockerOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) FlockerInput() *DeploymentSpecTemplateSpecVolumeFlocker {
	var returns *DeploymentSpecTemplateSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GcePersistentDisk() DeploymentSpecTemplateSpecVolumeGcePersistentDiskOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GcePersistentDiskInput() *DeploymentSpecTemplateSpecVolumeGcePersistentDisk {
	var returns *DeploymentSpecTemplateSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GitRepo() DeploymentSpecTemplateSpecVolumeGitRepoOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GitRepoInput() *DeploymentSpecTemplateSpecVolumeGitRepo {
	var returns *DeploymentSpecTemplateSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Glusterfs() DeploymentSpecTemplateSpecVolumeGlusterfsOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GlusterfsInput() *DeploymentSpecTemplateSpecVolumeGlusterfs {
	var returns *DeploymentSpecTemplateSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) HostPath() DeploymentSpecTemplateSpecVolumeHostPathOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) HostPathInput() *DeploymentSpecTemplateSpecVolumeHostPath {
	var returns *DeploymentSpecTemplateSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Iscsi() DeploymentSpecTemplateSpecVolumeIscsiOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) IscsiInput() *DeploymentSpecTemplateSpecVolumeIscsi {
	var returns *DeploymentSpecTemplateSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Local() DeploymentSpecTemplateSpecVolumeLocalOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) LocalInput() *DeploymentSpecTemplateSpecVolumeLocal {
	var returns *DeploymentSpecTemplateSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Nfs() DeploymentSpecTemplateSpecVolumeNfsOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) NfsInput() *DeploymentSpecTemplateSpecVolumeNfs {
	var returns *DeploymentSpecTemplateSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaim() DeploymentSpecTemplateSpecVolumePersistentVolumeClaimOutputReference {
	var returns DeploymentSpecTemplateSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PersistentVolumeClaimInput() *DeploymentSpecTemplateSpecVolumePersistentVolumeClaim {
	var returns *DeploymentSpecTemplateSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PhotonPersistentDisk() DeploymentSpecTemplateSpecVolumePhotonPersistentDiskOutputReference {
	var returns DeploymentSpecTemplateSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PhotonPersistentDiskInput() *DeploymentSpecTemplateSpecVolumePhotonPersistentDisk {
	var returns *DeploymentSpecTemplateSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Projected() DeploymentSpecTemplateSpecVolumeProjectedList {
	var returns DeploymentSpecTemplateSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Quobyte() DeploymentSpecTemplateSpecVolumeQuobyteOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) QuobyteInput() *DeploymentSpecTemplateSpecVolumeQuobyte {
	var returns *DeploymentSpecTemplateSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Rbd() DeploymentSpecTemplateSpecVolumeRbdOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) RbdInput() *DeploymentSpecTemplateSpecVolumeRbd {
	var returns *DeploymentSpecTemplateSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Secret() DeploymentSpecTemplateSpecVolumeSecretOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) SecretInput() *DeploymentSpecTemplateSpecVolumeSecret {
	var returns *DeploymentSpecTemplateSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) VsphereVolume() DeploymentSpecTemplateSpecVolumeVsphereVolumeOutputReference {
	var returns DeploymentSpecTemplateSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) VsphereVolumeInput() *DeploymentSpecTemplateSpecVolumeVsphereVolume {
	var returns *DeploymentSpecTemplateSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewDeploymentSpecTemplateSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DeploymentSpecTemplateSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewDeploymentSpecTemplateSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDeploymentSpecTemplateSpecVolumeOutputReference_Override(d DeploymentSpecTemplateSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.deployment.DeploymentSpecTemplateSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutAwsElasticBlockStore(value *DeploymentSpecTemplateSpecVolumeAwsElasticBlockStore) {
	if err := d.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutAzureDisk(value *DeploymentSpecTemplateSpecVolumeAzureDisk) {
	if err := d.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutAzureFile(value *DeploymentSpecTemplateSpecVolumeAzureFile) {
	if err := d.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutCephFs(value *DeploymentSpecTemplateSpecVolumeCephFs) {
	if err := d.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCephFs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutCinder(value *DeploymentSpecTemplateSpecVolumeCinder) {
	if err := d.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCinder",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutConfigMap(value *DeploymentSpecTemplateSpecVolumeConfigMap) {
	if err := d.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutCsi(value *DeploymentSpecTemplateSpecVolumeCsi) {
	if err := d.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutDownwardApi(value *DeploymentSpecTemplateSpecVolumeDownwardApi) {
	if err := d.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutEmptyDir(value *DeploymentSpecTemplateSpecVolumeEmptyDir) {
	if err := d.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutFc(value *DeploymentSpecTemplateSpecVolumeFc) {
	if err := d.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFc",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutFlexVolume(value *DeploymentSpecTemplateSpecVolumeFlexVolume) {
	if err := d.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutFlocker(value *DeploymentSpecTemplateSpecVolumeFlocker) {
	if err := d.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFlocker",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutGcePersistentDisk(value *DeploymentSpecTemplateSpecVolumeGcePersistentDisk) {
	if err := d.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutGitRepo(value *DeploymentSpecTemplateSpecVolumeGitRepo) {
	if err := d.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutGlusterfs(value *DeploymentSpecTemplateSpecVolumeGlusterfs) {
	if err := d.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutHostPath(value *DeploymentSpecTemplateSpecVolumeHostPath) {
	if err := d.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostPath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutIscsi(value *DeploymentSpecTemplateSpecVolumeIscsi) {
	if err := d.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIscsi",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutLocal(value *DeploymentSpecTemplateSpecVolumeLocal) {
	if err := d.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocal",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutNfs(value *DeploymentSpecTemplateSpecVolumeNfs) {
	if err := d.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNfs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutPersistentVolumeClaim(value *DeploymentSpecTemplateSpecVolumePersistentVolumeClaim) {
	if err := d.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutPhotonPersistentDisk(value *DeploymentSpecTemplateSpecVolumePhotonPersistentDisk) {
	if err := d.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := d.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProjected",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutQuobyte(value *DeploymentSpecTemplateSpecVolumeQuobyte) {
	if err := d.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutRbd(value *DeploymentSpecTemplateSpecVolumeRbd) {
	if err := d.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRbd",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutSecret(value *DeploymentSpecTemplateSpecVolumeSecret) {
	if err := d.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecret",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) PutVsphereVolume(value *DeploymentSpecTemplateSpecVolumeVsphereVolume) {
	if err := d.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		d,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		d,
		"resetCephFs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		d,
		"resetCinder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		d,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		d,
		"resetCsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		d,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		d,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		d,
		"resetFc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		d,
		"resetFlocker",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		d,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		d,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		d,
		"resetHostPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		d,
		"resetIscsi",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		d,
		"resetLocal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		d,
		"resetNfs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		d,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		d,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		d,
		"resetProjected",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		d,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		d,
		"resetRbd",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		d,
		"resetSecret",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DeploymentSpecTemplateSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

