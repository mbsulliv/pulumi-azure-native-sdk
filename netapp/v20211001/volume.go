// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20211001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Volume resource
type Volume struct {
	pulumi.CustomResourceState

	// Specifies whether the volume is enabled for Azure VMware Solution (AVS) datastore purpose
	AvsDataStore pulumi.StringPtrOutput `pulumi:"avsDataStore"`
	// UUID v4 or resource identifier used to identify the Backup.
	BackupId pulumi.StringPtrOutput `pulumi:"backupId"`
	// Unique Baremetal Tenant Identifier.
	BaremetalTenantId pulumi.StringOutput `pulumi:"baremetalTenantId"`
	// Pool Resource Id used in case of creating a volume through volume group
	CapacityPoolResourceId pulumi.StringPtrOutput `pulumi:"capacityPoolResourceId"`
	// When a volume is being restored from another volume's snapshot, will show the percentage completion of this cloning process. When this value is empty/null there is no cloning process currently happening on this volume. This value will update every 5 minutes during cloning.
	CloneProgress pulumi.IntOutput `pulumi:"cloneProgress"`
	// Specifies whether Cool Access(tiering) is enabled for the volume.
	CoolAccess pulumi.BoolPtrOutput `pulumi:"coolAccess"`
	// Specifies the number of days after which data that is not accessed by clients will be tiered.
	CoolnessPeriod pulumi.IntPtrOutput `pulumi:"coolnessPeriod"`
	// A unique file path for the volume. Used when creating mount targets
	CreationToken pulumi.StringOutput `pulumi:"creationToken"`
	// DataProtection type volumes include an object containing details of the replication
	DataProtection VolumePropertiesResponseDataProtectionPtrOutput `pulumi:"dataProtection"`
	// Default group quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies.
	DefaultGroupQuotaInKiBs pulumi.Float64PtrOutput `pulumi:"defaultGroupQuotaInKiBs"`
	// Default user quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies .
	DefaultUserQuotaInKiBs pulumi.Float64PtrOutput `pulumi:"defaultUserQuotaInKiBs"`
	// Flag indicating whether subvolume operations are enabled on the volume
	EnableSubvolumes pulumi.StringPtrOutput `pulumi:"enableSubvolumes"`
	// Encryption Key Source. Possible values are: 'Microsoft.NetApp'
	EncryptionKeySource pulumi.StringPtrOutput `pulumi:"encryptionKeySource"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Set of export policy rules
	ExportPolicy VolumePropertiesResponseExportPolicyPtrOutput `pulumi:"exportPolicy"`
	// Unique FileSystem Identifier.
	FileSystemId pulumi.StringOutput `pulumi:"fileSystemId"`
	// Specifies if default quota is enabled for the volume.
	IsDefaultQuotaEnabled pulumi.BoolPtrOutput `pulumi:"isDefaultQuotaEnabled"`
	// Restoring
	IsRestoring pulumi.BoolPtrOutput `pulumi:"isRestoring"`
	// Describe if a volume is KerberosEnabled. To be use with swagger version 2020-05-01 or later
	KerberosEnabled pulumi.BoolPtrOutput `pulumi:"kerberosEnabled"`
	// Specifies whether LDAP is enabled or not for a given NFS volume.
	LdapEnabled pulumi.BoolPtrOutput `pulumi:"ldapEnabled"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum number of files allowed. Needs a service request in order to be changed. Only allowed to be changed if volume quota is more than 4TiB.
	MaximumNumberOfFiles pulumi.Float64Output `pulumi:"maximumNumberOfFiles"`
	// List of mount targets
	MountTargets MountTargetPropertiesResponseArrayOutput `pulumi:"mountTargets"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Basic network, or Standard features available to the volume.
	NetworkFeatures pulumi.StringPtrOutput `pulumi:"networkFeatures"`
	// Network Sibling Set ID for the the group of volumes sharing networking resources.
	NetworkSiblingSetId pulumi.StringOutput `pulumi:"networkSiblingSetId"`
	// Application specific placement rules for the particular volume
	PlacementRules PlacementKeyValuePairsResponseArrayOutput `pulumi:"placementRules"`
	// Set of protocol types, default NFSv3, CIFS for SMB protocol
	ProtocolTypes pulumi.StringArrayOutput `pulumi:"protocolTypes"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Proximity placement group associated with the volume
	ProximityPlacementGroup pulumi.StringPtrOutput `pulumi:"proximityPlacementGroup"`
	// The security style of volume, default unix, defaults to ntfs for dual protocol or CIFS protocol
	SecurityStyle pulumi.StringPtrOutput `pulumi:"securityStyle"`
	// The service level of the file system
	ServiceLevel pulumi.StringPtrOutput `pulumi:"serviceLevel"`
	// Enables continuously available share property for smb volume. Only applicable for SMB volume
	SmbContinuouslyAvailable pulumi.BoolPtrOutput `pulumi:"smbContinuouslyAvailable"`
	// Enables encryption for in-flight smb3 data. Only applicable for SMB/DualProtocol volume. To be used with swagger version 2020-08-01 or later
	SmbEncryption pulumi.BoolPtrOutput `pulumi:"smbEncryption"`
	// If enabled (true) the volume will contain a read-only snapshot directory which provides access to each of the volume's snapshots (default to true).
	SnapshotDirectoryVisible pulumi.BoolPtrOutput `pulumi:"snapshotDirectoryVisible"`
	// UUID v4 or resource identifier used to identify the Snapshot.
	SnapshotId pulumi.StringPtrOutput `pulumi:"snapshotId"`
	// Provides storage to network proximity information for the volume.
	StorageToNetworkProximity pulumi.StringOutput `pulumi:"storageToNetworkProximity"`
	// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// T2 network information
	T2Network pulumi.StringOutput `pulumi:"t2Network"`
	// Resource tags.
	Tags            pulumi.StringMapOutput  `pulumi:"tags"`
	ThroughputMibps pulumi.Float64PtrOutput `pulumi:"throughputMibps"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// UNIX permissions for NFS volume accepted in octal 4 digit format. First digit selects the set user ID(4), set group ID (2) and sticky (1) attributes. Second digit selects permission for the owner of the file: read (4), write (2) and execute (1). Third selects permissions for other users in the same group. the fourth for other users not in the group. 0755 - gives read/write/execute permissions to owner and read/execute to group and other users.
	UnixPermissions pulumi.StringPtrOutput `pulumi:"unixPermissions"`
	// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
	UsageThreshold pulumi.Float64Output `pulumi:"usageThreshold"`
	// Volume Group Name
	VolumeGroupName pulumi.StringOutput `pulumi:"volumeGroupName"`
	// Volume spec name is the application specific designation or identifier for the particular volume in a volume group for e.g. data, log
	VolumeSpecName pulumi.StringPtrOutput `pulumi:"volumeSpecName"`
	// What type of volume is this. For destination volumes in Cross Region Replication, set type to DataProtection
	VolumeType pulumi.StringPtrOutput `pulumi:"volumeType"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.CreationToken == nil {
		return nil, errors.New("invalid value for required argument 'CreationToken'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.AvsDataStore == nil {
		args.AvsDataStore = pulumi.StringPtr("Disabled")
	}
	if args.CoolAccess == nil {
		args.CoolAccess = pulumi.BoolPtr(false)
	}
	if args.DefaultGroupQuotaInKiBs == nil {
		args.DefaultGroupQuotaInKiBs = pulumi.Float64Ptr(0.0)
	}
	if args.DefaultUserQuotaInKiBs == nil {
		args.DefaultUserQuotaInKiBs = pulumi.Float64Ptr(0.0)
	}
	if args.EnableSubvolumes == nil {
		args.EnableSubvolumes = pulumi.StringPtr("Disabled")
	}
	if args.IsDefaultQuotaEnabled == nil {
		args.IsDefaultQuotaEnabled = pulumi.BoolPtr(false)
	}
	if args.KerberosEnabled == nil {
		args.KerberosEnabled = pulumi.BoolPtr(false)
	}
	if args.LdapEnabled == nil {
		args.LdapEnabled = pulumi.BoolPtr(false)
	}
	if args.NetworkFeatures == nil {
		args.NetworkFeatures = pulumi.StringPtr("Basic")
	}
	if args.SecurityStyle == nil {
		args.SecurityStyle = pulumi.StringPtr("unix")
	}
	if args.SmbContinuouslyAvailable == nil {
		args.SmbContinuouslyAvailable = pulumi.BoolPtr(false)
	}
	if args.SmbEncryption == nil {
		args.SmbEncryption = pulumi.BoolPtr(false)
	}
	if args.SnapshotDirectoryVisible == nil {
		args.SnapshotDirectoryVisible = pulumi.BoolPtr(true)
	}
	if args.UnixPermissions == nil {
		args.UnixPermissions = pulumi.StringPtr("0770")
	}
	if args.UsageThreshold == nil {
		args.UsageThreshold = pulumi.Float64(107374182400.0)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101preview:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230501:Volume"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("azure-native:netapp/v20211001:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:netapp/v20211001:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
}

type VolumeState struct {
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// Specifies whether the volume is enabled for Azure VMware Solution (AVS) datastore purpose
	AvsDataStore *string `pulumi:"avsDataStore"`
	// UUID v4 or resource identifier used to identify the Backup.
	BackupId *string `pulumi:"backupId"`
	// Pool Resource Id used in case of creating a volume through volume group
	CapacityPoolResourceId *string `pulumi:"capacityPoolResourceId"`
	// Specifies whether Cool Access(tiering) is enabled for the volume.
	CoolAccess *bool `pulumi:"coolAccess"`
	// Specifies the number of days after which data that is not accessed by clients will be tiered.
	CoolnessPeriod *int `pulumi:"coolnessPeriod"`
	// A unique file path for the volume. Used when creating mount targets
	CreationToken string `pulumi:"creationToken"`
	// DataProtection type volumes include an object containing details of the replication
	DataProtection *VolumePropertiesDataProtection `pulumi:"dataProtection"`
	// Default group quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies.
	DefaultGroupQuotaInKiBs *float64 `pulumi:"defaultGroupQuotaInKiBs"`
	// Default user quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies .
	DefaultUserQuotaInKiBs *float64 `pulumi:"defaultUserQuotaInKiBs"`
	// Flag indicating whether subvolume operations are enabled on the volume
	EnableSubvolumes *string `pulumi:"enableSubvolumes"`
	// Encryption Key Source. Possible values are: 'Microsoft.NetApp'
	EncryptionKeySource *string `pulumi:"encryptionKeySource"`
	// Set of export policy rules
	ExportPolicy *VolumePropertiesExportPolicy `pulumi:"exportPolicy"`
	// Specifies if default quota is enabled for the volume.
	IsDefaultQuotaEnabled *bool `pulumi:"isDefaultQuotaEnabled"`
	// Restoring
	IsRestoring *bool `pulumi:"isRestoring"`
	// Describe if a volume is KerberosEnabled. To be use with swagger version 2020-05-01 or later
	KerberosEnabled *bool `pulumi:"kerberosEnabled"`
	// Specifies whether LDAP is enabled or not for a given NFS volume.
	LdapEnabled *bool `pulumi:"ldapEnabled"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Basic network, or Standard features available to the volume.
	NetworkFeatures *string `pulumi:"networkFeatures"`
	// Application specific placement rules for the particular volume
	PlacementRules []PlacementKeyValuePairs `pulumi:"placementRules"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// Set of protocol types, default NFSv3, CIFS for SMB protocol
	ProtocolTypes []string `pulumi:"protocolTypes"`
	// Proximity placement group associated with the volume
	ProximityPlacementGroup *string `pulumi:"proximityPlacementGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The security style of volume, default unix, defaults to ntfs for dual protocol or CIFS protocol
	SecurityStyle *string `pulumi:"securityStyle"`
	// The service level of the file system
	ServiceLevel *string `pulumi:"serviceLevel"`
	// Enables continuously available share property for smb volume. Only applicable for SMB volume
	SmbContinuouslyAvailable *bool `pulumi:"smbContinuouslyAvailable"`
	// Enables encryption for in-flight smb3 data. Only applicable for SMB/DualProtocol volume. To be used with swagger version 2020-08-01 or later
	SmbEncryption *bool `pulumi:"smbEncryption"`
	// If enabled (true) the volume will contain a read-only snapshot directory which provides access to each of the volume's snapshots (default to true).
	SnapshotDirectoryVisible *bool `pulumi:"snapshotDirectoryVisible"`
	// UUID v4 or resource identifier used to identify the Snapshot.
	SnapshotId *string `pulumi:"snapshotId"`
	// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
	SubnetId string `pulumi:"subnetId"`
	// Resource tags.
	Tags            map[string]string `pulumi:"tags"`
	ThroughputMibps *float64          `pulumi:"throughputMibps"`
	// UNIX permissions for NFS volume accepted in octal 4 digit format. First digit selects the set user ID(4), set group ID (2) and sticky (1) attributes. Second digit selects permission for the owner of the file: read (4), write (2) and execute (1). Third selects permissions for other users in the same group. the fourth for other users not in the group. 0755 - gives read/write/execute permissions to owner and read/execute to group and other users.
	UnixPermissions *string `pulumi:"unixPermissions"`
	// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
	UsageThreshold float64 `pulumi:"usageThreshold"`
	// The name of the volume
	VolumeName *string `pulumi:"volumeName"`
	// Volume spec name is the application specific designation or identifier for the particular volume in a volume group for e.g. data, log
	VolumeSpecName *string `pulumi:"volumeSpecName"`
	// What type of volume is this. For destination volumes in Cross Region Replication, set type to DataProtection
	VolumeType *string `pulumi:"volumeType"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// Specifies whether the volume is enabled for Azure VMware Solution (AVS) datastore purpose
	AvsDataStore pulumi.StringPtrInput
	// UUID v4 or resource identifier used to identify the Backup.
	BackupId pulumi.StringPtrInput
	// Pool Resource Id used in case of creating a volume through volume group
	CapacityPoolResourceId pulumi.StringPtrInput
	// Specifies whether Cool Access(tiering) is enabled for the volume.
	CoolAccess pulumi.BoolPtrInput
	// Specifies the number of days after which data that is not accessed by clients will be tiered.
	CoolnessPeriod pulumi.IntPtrInput
	// A unique file path for the volume. Used when creating mount targets
	CreationToken pulumi.StringInput
	// DataProtection type volumes include an object containing details of the replication
	DataProtection VolumePropertiesDataProtectionPtrInput
	// Default group quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies.
	DefaultGroupQuotaInKiBs pulumi.Float64PtrInput
	// Default user quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies .
	DefaultUserQuotaInKiBs pulumi.Float64PtrInput
	// Flag indicating whether subvolume operations are enabled on the volume
	EnableSubvolumes pulumi.StringPtrInput
	// Encryption Key Source. Possible values are: 'Microsoft.NetApp'
	EncryptionKeySource pulumi.StringPtrInput
	// Set of export policy rules
	ExportPolicy VolumePropertiesExportPolicyPtrInput
	// Specifies if default quota is enabled for the volume.
	IsDefaultQuotaEnabled pulumi.BoolPtrInput
	// Restoring
	IsRestoring pulumi.BoolPtrInput
	// Describe if a volume is KerberosEnabled. To be use with swagger version 2020-05-01 or later
	KerberosEnabled pulumi.BoolPtrInput
	// Specifies whether LDAP is enabled or not for a given NFS volume.
	LdapEnabled pulumi.BoolPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Basic network, or Standard features available to the volume.
	NetworkFeatures pulumi.StringPtrInput
	// Application specific placement rules for the particular volume
	PlacementRules PlacementKeyValuePairsArrayInput
	// The name of the capacity pool
	PoolName pulumi.StringInput
	// Set of protocol types, default NFSv3, CIFS for SMB protocol
	ProtocolTypes pulumi.StringArrayInput
	// Proximity placement group associated with the volume
	ProximityPlacementGroup pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The security style of volume, default unix, defaults to ntfs for dual protocol or CIFS protocol
	SecurityStyle pulumi.StringPtrInput
	// The service level of the file system
	ServiceLevel pulumi.StringPtrInput
	// Enables continuously available share property for smb volume. Only applicable for SMB volume
	SmbContinuouslyAvailable pulumi.BoolPtrInput
	// Enables encryption for in-flight smb3 data. Only applicable for SMB/DualProtocol volume. To be used with swagger version 2020-08-01 or later
	SmbEncryption pulumi.BoolPtrInput
	// If enabled (true) the volume will contain a read-only snapshot directory which provides access to each of the volume's snapshots (default to true).
	SnapshotDirectoryVisible pulumi.BoolPtrInput
	// UUID v4 or resource identifier used to identify the Snapshot.
	SnapshotId pulumi.StringPtrInput
	// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
	SubnetId pulumi.StringInput
	// Resource tags.
	Tags            pulumi.StringMapInput
	ThroughputMibps pulumi.Float64PtrInput
	// UNIX permissions for NFS volume accepted in octal 4 digit format. First digit selects the set user ID(4), set group ID (2) and sticky (1) attributes. Second digit selects permission for the owner of the file: read (4), write (2) and execute (1). Third selects permissions for other users in the same group. the fourth for other users not in the group. 0755 - gives read/write/execute permissions to owner and read/execute to group and other users.
	UnixPermissions pulumi.StringPtrInput
	// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
	UsageThreshold pulumi.Float64Input
	// The name of the volume
	VolumeName pulumi.StringPtrInput
	// Volume spec name is the application specific designation or identifier for the particular volume in a volume group for e.g. data, log
	VolumeSpecName pulumi.StringPtrInput
	// What type of volume is this. For destination volumes in Cross Region Replication, set type to DataProtection
	VolumeType pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

// Specifies whether the volume is enabled for Azure VMware Solution (AVS) datastore purpose
func (o VolumeOutput) AvsDataStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.AvsDataStore }).(pulumi.StringPtrOutput)
}

// UUID v4 or resource identifier used to identify the Backup.
func (o VolumeOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.BackupId }).(pulumi.StringPtrOutput)
}

// Unique Baremetal Tenant Identifier.
func (o VolumeOutput) BaremetalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.BaremetalTenantId }).(pulumi.StringOutput)
}

// Pool Resource Id used in case of creating a volume through volume group
func (o VolumeOutput) CapacityPoolResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.CapacityPoolResourceId }).(pulumi.StringPtrOutput)
}

// When a volume is being restored from another volume's snapshot, will show the percentage completion of this cloning process. When this value is empty/null there is no cloning process currently happening on this volume. This value will update every 5 minutes during cloning.
func (o VolumeOutput) CloneProgress() pulumi.IntOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntOutput { return v.CloneProgress }).(pulumi.IntOutput)
}

// Specifies whether Cool Access(tiering) is enabled for the volume.
func (o VolumeOutput) CoolAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.CoolAccess }).(pulumi.BoolPtrOutput)
}

// Specifies the number of days after which data that is not accessed by clients will be tiered.
func (o VolumeOutput) CoolnessPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.CoolnessPeriod }).(pulumi.IntPtrOutput)
}

// A unique file path for the volume. Used when creating mount targets
func (o VolumeOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.CreationToken }).(pulumi.StringOutput)
}

// DataProtection type volumes include an object containing details of the replication
func (o VolumeOutput) DataProtection() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyT(func(v *Volume) VolumePropertiesResponseDataProtectionPtrOutput { return v.DataProtection }).(VolumePropertiesResponseDataProtectionPtrOutput)
}

// Default group quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies.
func (o VolumeOutput) DefaultGroupQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.DefaultGroupQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

// Default user quota for volume in KiBs. If isDefaultQuotaEnabled is set, the minimum value of 4 KiBs applies .
func (o VolumeOutput) DefaultUserQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.DefaultUserQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

// Flag indicating whether subvolume operations are enabled on the volume
func (o VolumeOutput) EnableSubvolumes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.EnableSubvolumes }).(pulumi.StringPtrOutput)
}

// Encryption Key Source. Possible values are: 'Microsoft.NetApp'
func (o VolumeOutput) EncryptionKeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.EncryptionKeySource }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o VolumeOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Set of export policy rules
func (o VolumeOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v *Volume) VolumePropertiesResponseExportPolicyPtrOutput { return v.ExportPolicy }).(VolumePropertiesResponseExportPolicyPtrOutput)
}

// Unique FileSystem Identifier.
func (o VolumeOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

// Specifies if default quota is enabled for the volume.
func (o VolumeOutput) IsDefaultQuotaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.IsDefaultQuotaEnabled }).(pulumi.BoolPtrOutput)
}

// Restoring
func (o VolumeOutput) IsRestoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.IsRestoring }).(pulumi.BoolPtrOutput)
}

// Describe if a volume is KerberosEnabled. To be use with swagger version 2020-05-01 or later
func (o VolumeOutput) KerberosEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.KerberosEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies whether LDAP is enabled or not for a given NFS volume.
func (o VolumeOutput) LdapEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.LdapEnabled }).(pulumi.BoolPtrOutput)
}

// The geo-location where the resource lives
func (o VolumeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maximum number of files allowed. Needs a service request in order to be changed. Only allowed to be changed if volume quota is more than 4TiB.
func (o VolumeOutput) MaximumNumberOfFiles() pulumi.Float64Output {
	return o.ApplyT(func(v *Volume) pulumi.Float64Output { return v.MaximumNumberOfFiles }).(pulumi.Float64Output)
}

// List of mount targets
func (o VolumeOutput) MountTargets() MountTargetPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *Volume) MountTargetPropertiesResponseArrayOutput { return v.MountTargets }).(MountTargetPropertiesResponseArrayOutput)
}

// The name of the resource
func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Basic network, or Standard features available to the volume.
func (o VolumeOutput) NetworkFeatures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.NetworkFeatures }).(pulumi.StringPtrOutput)
}

// Network Sibling Set ID for the the group of volumes sharing networking resources.
func (o VolumeOutput) NetworkSiblingSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.NetworkSiblingSetId }).(pulumi.StringOutput)
}

// Application specific placement rules for the particular volume
func (o VolumeOutput) PlacementRules() PlacementKeyValuePairsResponseArrayOutput {
	return o.ApplyT(func(v *Volume) PlacementKeyValuePairsResponseArrayOutput { return v.PlacementRules }).(PlacementKeyValuePairsResponseArrayOutput)
}

// Set of protocol types, default NFSv3, CIFS for SMB protocol
func (o VolumeOutput) ProtocolTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringArrayOutput { return v.ProtocolTypes }).(pulumi.StringArrayOutput)
}

// Azure lifecycle management
func (o VolumeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Proximity placement group associated with the volume
func (o VolumeOutput) ProximityPlacementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.ProximityPlacementGroup }).(pulumi.StringPtrOutput)
}

// The security style of volume, default unix, defaults to ntfs for dual protocol or CIFS protocol
func (o VolumeOutput) SecurityStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SecurityStyle }).(pulumi.StringPtrOutput)
}

// The service level of the file system
func (o VolumeOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

// Enables continuously available share property for smb volume. Only applicable for SMB volume
func (o VolumeOutput) SmbContinuouslyAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.SmbContinuouslyAvailable }).(pulumi.BoolPtrOutput)
}

// Enables encryption for in-flight smb3 data. Only applicable for SMB/DualProtocol volume. To be used with swagger version 2020-08-01 or later
func (o VolumeOutput) SmbEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.SmbEncryption }).(pulumi.BoolPtrOutput)
}

// If enabled (true) the volume will contain a read-only snapshot directory which provides access to each of the volume's snapshots (default to true).
func (o VolumeOutput) SnapshotDirectoryVisible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.SnapshotDirectoryVisible }).(pulumi.BoolPtrOutput)
}

// UUID v4 or resource identifier used to identify the Snapshot.
func (o VolumeOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

// Provides storage to network proximity information for the volume.
func (o VolumeOutput) StorageToNetworkProximity() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.StorageToNetworkProximity }).(pulumi.StringOutput)
}

// The Azure Resource URI for a delegated subnet. Must have the delegation Microsoft.NetApp/volumes
func (o VolumeOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VolumeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Volume) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// T2 network information
func (o VolumeOutput) T2Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.T2Network }).(pulumi.StringOutput)
}

// Resource tags.
func (o VolumeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VolumeOutput) ThroughputMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.ThroughputMibps }).(pulumi.Float64PtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// UNIX permissions for NFS volume accepted in octal 4 digit format. First digit selects the set user ID(4), set group ID (2) and sticky (1) attributes. Second digit selects permission for the owner of the file: read (4), write (2) and execute (1). Third selects permissions for other users in the same group. the fourth for other users not in the group. 0755 - gives read/write/execute permissions to owner and read/execute to group and other users.
func (o VolumeOutput) UnixPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.UnixPermissions }).(pulumi.StringPtrOutput)
}

// Maximum storage quota allowed for a file system in bytes. This is a soft quota used for alerting only. Minimum size is 100 GiB. Upper limit is 100TiB. Specified in bytes.
func (o VolumeOutput) UsageThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v *Volume) pulumi.Float64Output { return v.UsageThreshold }).(pulumi.Float64Output)
}

// Volume Group Name
func (o VolumeOutput) VolumeGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeGroupName }).(pulumi.StringOutput)
}

// Volume spec name is the application specific designation or identifier for the particular volume in a volume group for e.g. data, log
func (o VolumeOutput) VolumeSpecName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.VolumeSpecName }).(pulumi.StringPtrOutput)
}

// What type of volume is this. For destination volumes in Cross Region Replication, set type to DataProtection
func (o VolumeOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
