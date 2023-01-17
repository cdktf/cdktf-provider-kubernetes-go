package pod

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/jsii"

	"github.com/cdktf/cdktf-provider-kubernetes-go/kubernetes/v5/pod/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PodSpecVolumeOutputReference interface {
	cdktf.ComplexObject
	AwsElasticBlockStore() PodSpecVolumeAwsElasticBlockStoreOutputReference
	AwsElasticBlockStoreInput() *PodSpecVolumeAwsElasticBlockStore
	AzureDisk() PodSpecVolumeAzureDiskOutputReference
	AzureDiskInput() *PodSpecVolumeAzureDisk
	AzureFile() PodSpecVolumeAzureFileOutputReference
	AzureFileInput() *PodSpecVolumeAzureFile
	CephFs() PodSpecVolumeCephFsOutputReference
	CephFsInput() *PodSpecVolumeCephFs
	Cinder() PodSpecVolumeCinderOutputReference
	CinderInput() *PodSpecVolumeCinder
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
	ConfigMap() PodSpecVolumeConfigMapOutputReference
	ConfigMapInput() *PodSpecVolumeConfigMap
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Csi() PodSpecVolumeCsiOutputReference
	CsiInput() *PodSpecVolumeCsi
	DownwardApi() PodSpecVolumeDownwardApiOutputReference
	DownwardApiInput() *PodSpecVolumeDownwardApi
	EmptyDir() PodSpecVolumeEmptyDirOutputReference
	EmptyDirInput() *PodSpecVolumeEmptyDir
	Fc() PodSpecVolumeFcOutputReference
	FcInput() *PodSpecVolumeFc
	FlexVolume() PodSpecVolumeFlexVolumeOutputReference
	FlexVolumeInput() *PodSpecVolumeFlexVolume
	Flocker() PodSpecVolumeFlockerOutputReference
	FlockerInput() *PodSpecVolumeFlocker
	// Experimental.
	Fqn() *string
	GcePersistentDisk() PodSpecVolumeGcePersistentDiskOutputReference
	GcePersistentDiskInput() *PodSpecVolumeGcePersistentDisk
	GitRepo() PodSpecVolumeGitRepoOutputReference
	GitRepoInput() *PodSpecVolumeGitRepo
	Glusterfs() PodSpecVolumeGlusterfsOutputReference
	GlusterfsInput() *PodSpecVolumeGlusterfs
	HostPath() PodSpecVolumeHostPathOutputReference
	HostPathInput() *PodSpecVolumeHostPath
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Iscsi() PodSpecVolumeIscsiOutputReference
	IscsiInput() *PodSpecVolumeIscsi
	Local() PodSpecVolumeLocalOutputReference
	LocalInput() *PodSpecVolumeLocal
	Name() *string
	SetName(val *string)
	NameInput() *string
	Nfs() PodSpecVolumeNfsOutputReference
	NfsInput() *PodSpecVolumeNfs
	PersistentVolumeClaim() PodSpecVolumePersistentVolumeClaimOutputReference
	PersistentVolumeClaimInput() *PodSpecVolumePersistentVolumeClaim
	PhotonPersistentDisk() PodSpecVolumePhotonPersistentDiskOutputReference
	PhotonPersistentDiskInput() *PodSpecVolumePhotonPersistentDisk
	Projected() PodSpecVolumeProjectedList
	ProjectedInput() interface{}
	Quobyte() PodSpecVolumeQuobyteOutputReference
	QuobyteInput() *PodSpecVolumeQuobyte
	Rbd() PodSpecVolumeRbdOutputReference
	RbdInput() *PodSpecVolumeRbd
	Secret() PodSpecVolumeSecretOutputReference
	SecretInput() *PodSpecVolumeSecret
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VsphereVolume() PodSpecVolumeVsphereVolumeOutputReference
	VsphereVolumeInput() *PodSpecVolumeVsphereVolume
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
	PutAwsElasticBlockStore(value *PodSpecVolumeAwsElasticBlockStore)
	PutAzureDisk(value *PodSpecVolumeAzureDisk)
	PutAzureFile(value *PodSpecVolumeAzureFile)
	PutCephFs(value *PodSpecVolumeCephFs)
	PutCinder(value *PodSpecVolumeCinder)
	PutConfigMap(value *PodSpecVolumeConfigMap)
	PutCsi(value *PodSpecVolumeCsi)
	PutDownwardApi(value *PodSpecVolumeDownwardApi)
	PutEmptyDir(value *PodSpecVolumeEmptyDir)
	PutFc(value *PodSpecVolumeFc)
	PutFlexVolume(value *PodSpecVolumeFlexVolume)
	PutFlocker(value *PodSpecVolumeFlocker)
	PutGcePersistentDisk(value *PodSpecVolumeGcePersistentDisk)
	PutGitRepo(value *PodSpecVolumeGitRepo)
	PutGlusterfs(value *PodSpecVolumeGlusterfs)
	PutHostPath(value *PodSpecVolumeHostPath)
	PutIscsi(value *PodSpecVolumeIscsi)
	PutLocal(value *PodSpecVolumeLocal)
	PutNfs(value *PodSpecVolumeNfs)
	PutPersistentVolumeClaim(value *PodSpecVolumePersistentVolumeClaim)
	PutPhotonPersistentDisk(value *PodSpecVolumePhotonPersistentDisk)
	PutProjected(value interface{})
	PutQuobyte(value *PodSpecVolumeQuobyte)
	PutRbd(value *PodSpecVolumeRbd)
	PutSecret(value *PodSpecVolumeSecret)
	PutVsphereVolume(value *PodSpecVolumeVsphereVolume)
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

// The jsii proxy struct for PodSpecVolumeOutputReference
type jsiiProxy_PodSpecVolumeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) AwsElasticBlockStore() PodSpecVolumeAwsElasticBlockStoreOutputReference {
	var returns PodSpecVolumeAwsElasticBlockStoreOutputReference
	_jsii_.Get(
		j,
		"awsElasticBlockStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) AwsElasticBlockStoreInput() *PodSpecVolumeAwsElasticBlockStore {
	var returns *PodSpecVolumeAwsElasticBlockStore
	_jsii_.Get(
		j,
		"awsElasticBlockStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) AzureDisk() PodSpecVolumeAzureDiskOutputReference {
	var returns PodSpecVolumeAzureDiskOutputReference
	_jsii_.Get(
		j,
		"azureDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) AzureDiskInput() *PodSpecVolumeAzureDisk {
	var returns *PodSpecVolumeAzureDisk
	_jsii_.Get(
		j,
		"azureDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) AzureFile() PodSpecVolumeAzureFileOutputReference {
	var returns PodSpecVolumeAzureFileOutputReference
	_jsii_.Get(
		j,
		"azureFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) AzureFileInput() *PodSpecVolumeAzureFile {
	var returns *PodSpecVolumeAzureFile
	_jsii_.Get(
		j,
		"azureFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) CephFs() PodSpecVolumeCephFsOutputReference {
	var returns PodSpecVolumeCephFsOutputReference
	_jsii_.Get(
		j,
		"cephFs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) CephFsInput() *PodSpecVolumeCephFs {
	var returns *PodSpecVolumeCephFs
	_jsii_.Get(
		j,
		"cephFsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Cinder() PodSpecVolumeCinderOutputReference {
	var returns PodSpecVolumeCinderOutputReference
	_jsii_.Get(
		j,
		"cinder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) CinderInput() *PodSpecVolumeCinder {
	var returns *PodSpecVolumeCinder
	_jsii_.Get(
		j,
		"cinderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) ConfigMap() PodSpecVolumeConfigMapOutputReference {
	var returns PodSpecVolumeConfigMapOutputReference
	_jsii_.Get(
		j,
		"configMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) ConfigMapInput() *PodSpecVolumeConfigMap {
	var returns *PodSpecVolumeConfigMap
	_jsii_.Get(
		j,
		"configMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Csi() PodSpecVolumeCsiOutputReference {
	var returns PodSpecVolumeCsiOutputReference
	_jsii_.Get(
		j,
		"csi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) CsiInput() *PodSpecVolumeCsi {
	var returns *PodSpecVolumeCsi
	_jsii_.Get(
		j,
		"csiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) DownwardApi() PodSpecVolumeDownwardApiOutputReference {
	var returns PodSpecVolumeDownwardApiOutputReference
	_jsii_.Get(
		j,
		"downwardApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) DownwardApiInput() *PodSpecVolumeDownwardApi {
	var returns *PodSpecVolumeDownwardApi
	_jsii_.Get(
		j,
		"downwardApiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) EmptyDir() PodSpecVolumeEmptyDirOutputReference {
	var returns PodSpecVolumeEmptyDirOutputReference
	_jsii_.Get(
		j,
		"emptyDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) EmptyDirInput() *PodSpecVolumeEmptyDir {
	var returns *PodSpecVolumeEmptyDir
	_jsii_.Get(
		j,
		"emptyDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Fc() PodSpecVolumeFcOutputReference {
	var returns PodSpecVolumeFcOutputReference
	_jsii_.Get(
		j,
		"fc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) FcInput() *PodSpecVolumeFc {
	var returns *PodSpecVolumeFc
	_jsii_.Get(
		j,
		"fcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) FlexVolume() PodSpecVolumeFlexVolumeOutputReference {
	var returns PodSpecVolumeFlexVolumeOutputReference
	_jsii_.Get(
		j,
		"flexVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) FlexVolumeInput() *PodSpecVolumeFlexVolume {
	var returns *PodSpecVolumeFlexVolume
	_jsii_.Get(
		j,
		"flexVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Flocker() PodSpecVolumeFlockerOutputReference {
	var returns PodSpecVolumeFlockerOutputReference
	_jsii_.Get(
		j,
		"flocker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) FlockerInput() *PodSpecVolumeFlocker {
	var returns *PodSpecVolumeFlocker
	_jsii_.Get(
		j,
		"flockerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) GcePersistentDisk() PodSpecVolumeGcePersistentDiskOutputReference {
	var returns PodSpecVolumeGcePersistentDiskOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) GcePersistentDiskInput() *PodSpecVolumeGcePersistentDisk {
	var returns *PodSpecVolumeGcePersistentDisk
	_jsii_.Get(
		j,
		"gcePersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) GitRepo() PodSpecVolumeGitRepoOutputReference {
	var returns PodSpecVolumeGitRepoOutputReference
	_jsii_.Get(
		j,
		"gitRepo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) GitRepoInput() *PodSpecVolumeGitRepo {
	var returns *PodSpecVolumeGitRepo
	_jsii_.Get(
		j,
		"gitRepoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Glusterfs() PodSpecVolumeGlusterfsOutputReference {
	var returns PodSpecVolumeGlusterfsOutputReference
	_jsii_.Get(
		j,
		"glusterfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) GlusterfsInput() *PodSpecVolumeGlusterfs {
	var returns *PodSpecVolumeGlusterfs
	_jsii_.Get(
		j,
		"glusterfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) HostPath() PodSpecVolumeHostPathOutputReference {
	var returns PodSpecVolumeHostPathOutputReference
	_jsii_.Get(
		j,
		"hostPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) HostPathInput() *PodSpecVolumeHostPath {
	var returns *PodSpecVolumeHostPath
	_jsii_.Get(
		j,
		"hostPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Iscsi() PodSpecVolumeIscsiOutputReference {
	var returns PodSpecVolumeIscsiOutputReference
	_jsii_.Get(
		j,
		"iscsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) IscsiInput() *PodSpecVolumeIscsi {
	var returns *PodSpecVolumeIscsi
	_jsii_.Get(
		j,
		"iscsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Local() PodSpecVolumeLocalOutputReference {
	var returns PodSpecVolumeLocalOutputReference
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) LocalInput() *PodSpecVolumeLocal {
	var returns *PodSpecVolumeLocal
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Nfs() PodSpecVolumeNfsOutputReference {
	var returns PodSpecVolumeNfsOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) NfsInput() *PodSpecVolumeNfs {
	var returns *PodSpecVolumeNfs
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) PersistentVolumeClaim() PodSpecVolumePersistentVolumeClaimOutputReference {
	var returns PodSpecVolumePersistentVolumeClaimOutputReference
	_jsii_.Get(
		j,
		"persistentVolumeClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) PersistentVolumeClaimInput() *PodSpecVolumePersistentVolumeClaim {
	var returns *PodSpecVolumePersistentVolumeClaim
	_jsii_.Get(
		j,
		"persistentVolumeClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) PhotonPersistentDisk() PodSpecVolumePhotonPersistentDiskOutputReference {
	var returns PodSpecVolumePhotonPersistentDiskOutputReference
	_jsii_.Get(
		j,
		"photonPersistentDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) PhotonPersistentDiskInput() *PodSpecVolumePhotonPersistentDisk {
	var returns *PodSpecVolumePhotonPersistentDisk
	_jsii_.Get(
		j,
		"photonPersistentDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Projected() PodSpecVolumeProjectedList {
	var returns PodSpecVolumeProjectedList
	_jsii_.Get(
		j,
		"projected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) ProjectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Quobyte() PodSpecVolumeQuobyteOutputReference {
	var returns PodSpecVolumeQuobyteOutputReference
	_jsii_.Get(
		j,
		"quobyte",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) QuobyteInput() *PodSpecVolumeQuobyte {
	var returns *PodSpecVolumeQuobyte
	_jsii_.Get(
		j,
		"quobyteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Rbd() PodSpecVolumeRbdOutputReference {
	var returns PodSpecVolumeRbdOutputReference
	_jsii_.Get(
		j,
		"rbd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) RbdInput() *PodSpecVolumeRbd {
	var returns *PodSpecVolumeRbd
	_jsii_.Get(
		j,
		"rbdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) Secret() PodSpecVolumeSecretOutputReference {
	var returns PodSpecVolumeSecretOutputReference
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) SecretInput() *PodSpecVolumeSecret {
	var returns *PodSpecVolumeSecret
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) VsphereVolume() PodSpecVolumeVsphereVolumeOutputReference {
	var returns PodSpecVolumeVsphereVolumeOutputReference
	_jsii_.Get(
		j,
		"vsphereVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PodSpecVolumeOutputReference) VsphereVolumeInput() *PodSpecVolumeVsphereVolume {
	var returns *PodSpecVolumeVsphereVolume
	_jsii_.Get(
		j,
		"vsphereVolumeInput",
		&returns,
	)
	return returns
}


func NewPodSpecVolumeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PodSpecVolumeOutputReference {
	_init_.Initialize()

	if err := validateNewPodSpecVolumeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PodSpecVolumeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPodSpecVolumeOutputReference_Override(p PodSpecVolumeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-kubernetes.pod.PodSpecVolumeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PodSpecVolumeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PodSpecVolumeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutAwsElasticBlockStore(value *PodSpecVolumeAwsElasticBlockStore) {
	if err := p.validatePutAwsElasticBlockStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAwsElasticBlockStore",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutAzureDisk(value *PodSpecVolumeAzureDisk) {
	if err := p.validatePutAzureDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutAzureFile(value *PodSpecVolumeAzureFile) {
	if err := p.validatePutAzureFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureFile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutCephFs(value *PodSpecVolumeCephFs) {
	if err := p.validatePutCephFsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCephFs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutCinder(value *PodSpecVolumeCinder) {
	if err := p.validatePutCinderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCinder",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutConfigMap(value *PodSpecVolumeConfigMap) {
	if err := p.validatePutConfigMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConfigMap",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutCsi(value *PodSpecVolumeCsi) {
	if err := p.validatePutCsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCsi",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutDownwardApi(value *PodSpecVolumeDownwardApi) {
	if err := p.validatePutDownwardApiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDownwardApi",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutEmptyDir(value *PodSpecVolumeEmptyDir) {
	if err := p.validatePutEmptyDirParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEmptyDir",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutFc(value *PodSpecVolumeFc) {
	if err := p.validatePutFcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFc",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutFlexVolume(value *PodSpecVolumeFlexVolume) {
	if err := p.validatePutFlexVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFlexVolume",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutFlocker(value *PodSpecVolumeFlocker) {
	if err := p.validatePutFlockerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFlocker",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutGcePersistentDisk(value *PodSpecVolumeGcePersistentDisk) {
	if err := p.validatePutGcePersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGcePersistentDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutGitRepo(value *PodSpecVolumeGitRepo) {
	if err := p.validatePutGitRepoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGitRepo",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutGlusterfs(value *PodSpecVolumeGlusterfs) {
	if err := p.validatePutGlusterfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGlusterfs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutHostPath(value *PodSpecVolumeHostPath) {
	if err := p.validatePutHostPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHostPath",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutIscsi(value *PodSpecVolumeIscsi) {
	if err := p.validatePutIscsiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIscsi",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutLocal(value *PodSpecVolumeLocal) {
	if err := p.validatePutLocalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putLocal",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutNfs(value *PodSpecVolumeNfs) {
	if err := p.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNfs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutPersistentVolumeClaim(value *PodSpecVolumePersistentVolumeClaim) {
	if err := p.validatePutPersistentVolumeClaimParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPersistentVolumeClaim",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutPhotonPersistentDisk(value *PodSpecVolumePhotonPersistentDisk) {
	if err := p.validatePutPhotonPersistentDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPhotonPersistentDisk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutProjected(value interface{}) {
	if err := p.validatePutProjectedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putProjected",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutQuobyte(value *PodSpecVolumeQuobyte) {
	if err := p.validatePutQuobyteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putQuobyte",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutRbd(value *PodSpecVolumeRbd) {
	if err := p.validatePutRbdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRbd",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutSecret(value *PodSpecVolumeSecret) {
	if err := p.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSecret",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) PutVsphereVolume(value *PodSpecVolumeVsphereVolume) {
	if err := p.validatePutVsphereVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVsphereVolume",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetAwsElasticBlockStore() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsElasticBlockStore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetAzureDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetAzureFile() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureFile",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetCephFs() {
	_jsii_.InvokeVoid(
		p,
		"resetCephFs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetCinder() {
	_jsii_.InvokeVoid(
		p,
		"resetCinder",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetConfigMap() {
	_jsii_.InvokeVoid(
		p,
		"resetConfigMap",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetCsi() {
	_jsii_.InvokeVoid(
		p,
		"resetCsi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetDownwardApi() {
	_jsii_.InvokeVoid(
		p,
		"resetDownwardApi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetEmptyDir() {
	_jsii_.InvokeVoid(
		p,
		"resetEmptyDir",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetFc() {
	_jsii_.InvokeVoid(
		p,
		"resetFc",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetFlexVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetFlexVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetFlocker() {
	_jsii_.InvokeVoid(
		p,
		"resetFlocker",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetGcePersistentDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetGcePersistentDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetGitRepo() {
	_jsii_.InvokeVoid(
		p,
		"resetGitRepo",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetGlusterfs() {
	_jsii_.InvokeVoid(
		p,
		"resetGlusterfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetHostPath() {
	_jsii_.InvokeVoid(
		p,
		"resetHostPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetIscsi() {
	_jsii_.InvokeVoid(
		p,
		"resetIscsi",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetLocal() {
	_jsii_.InvokeVoid(
		p,
		"resetLocal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		p,
		"resetNfs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetPersistentVolumeClaim() {
	_jsii_.InvokeVoid(
		p,
		"resetPersistentVolumeClaim",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetPhotonPersistentDisk() {
	_jsii_.InvokeVoid(
		p,
		"resetPhotonPersistentDisk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetProjected() {
	_jsii_.InvokeVoid(
		p,
		"resetProjected",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetQuobyte() {
	_jsii_.InvokeVoid(
		p,
		"resetQuobyte",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetRbd() {
	_jsii_.InvokeVoid(
		p,
		"resetRbd",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		p,
		"resetSecret",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) ResetVsphereVolume() {
	_jsii_.InvokeVoid(
		p,
		"resetVsphereVolume",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PodSpecVolumeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PodSpecVolumeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

