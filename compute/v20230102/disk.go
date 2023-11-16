// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230102

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Disk resource.
type Disk struct {
	pulumi.CustomResourceState

	// Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is disabled by default. Does not apply to Ultra disks.
	BurstingEnabled pulumi.BoolPtrOutput `pulumi:"burstingEnabled"`
	// Latest time when bursting was last enabled on a disk.
	BurstingEnabledTime pulumi.StringOutput `pulumi:"burstingEnabledTime"`
	// Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent pulumi.Float64PtrOutput `pulumi:"completionPercent"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponseOutput `pulumi:"creationData"`
	// Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode pulumi.StringPtrOutput `pulumi:"dataAccessAuthMode"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId pulumi.StringPtrOutput `pulumi:"diskAccessId"`
	// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly pulumi.Float64PtrOutput `pulumi:"diskIOPSReadOnly"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadWrite pulumi.Float64PtrOutput `pulumi:"diskIOPSReadWrite"`
	// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly pulumi.Float64PtrOutput `pulumi:"diskMBpsReadOnly"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite pulumi.Float64PtrOutput `pulumi:"diskMBpsReadWrite"`
	// The size of the disk in bytes. This field is read only.
	DiskSizeBytes pulumi.Float64Output `pulumi:"diskSizeBytes"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrOutput `pulumi:"diskSizeGB"`
	// The state of the disk.
	DiskState pulumi.StringOutput `pulumi:"diskState"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption EncryptionResponsePtrOutput `pulumi:"encryption"`
	// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection EncryptionSettingsCollectionResponsePtrOutput `pulumi:"encryptionSettingsCollection"`
	// The extended location where the disk will be created. Extended location cannot be changed.
	ExtendedLocation ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrOutput `pulumi:"hyperVGeneration"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy pulumi.StringOutput `pulumi:"managedBy"`
	// List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
	ManagedByExtended pulumi.StringArrayOutput `pulumi:"managedByExtended"`
	// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
	MaxShares pulumi.IntPtrOutput `pulumi:"maxShares"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy pulumi.StringPtrOutput `pulumi:"networkAccessPolicy"`
	// Setting this property to true improves reliability and performance of data disks that are frequently (more than 5 times a day) by detached from one virtual machine and attached to another. This property should not be set for disks that are not detached and attached frequently as it causes the disks to not align with the fault domain of the virtual machine.
	OptimizedForFrequentAttach pulumi.BoolPtrOutput `pulumi:"optimizedForFrequentAttach"`
	// The Operating System type.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// Properties of the disk for which update is pending.
	PropertyUpdatesInProgress PropertyUpdatesInProgressResponseOutput `pulumi:"propertyUpdatesInProgress"`
	// The disk provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Policy for controlling export on the disk.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Purchase plan information for the the image from which the OS disk was created. E.g. - {name: 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
	PurchasePlan PurchasePlanResponsePtrOutput `pulumi:"purchasePlan"`
	// Contains the security related information for the resource.
	SecurityProfile DiskSecurityProfileResponsePtrOutput `pulumi:"securityProfile"`
	// Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
	ShareInfo ShareInfoElementResponseArrayOutput `pulumi:"shareInfo"`
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS, or PremiumV2_LRS.
	Sku DiskSkuResponsePtrOutput `pulumi:"sku"`
	// List of supported capabilities for the image from which the OS disk was created.
	SupportedCapabilities SupportedCapabilitiesResponsePtrOutput `pulumi:"supportedCapabilities"`
	// Indicates the OS on a disk supports hibernation.
	SupportsHibernation pulumi.BoolPtrOutput `pulumi:"supportsHibernation"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Performance tier of the disk (e.g, P4, S10) as described here: https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
	// The time when the disk was created.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Unique Guid identifying the resource.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// The Logical zone list for Disk.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewDisk registers a new resource with the given unique name, arguments, and options.
func NewDisk(ctx *pulumi.Context,
	name string, args *DiskArgs, opts ...pulumi.ResourceOption) (*Disk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CreationData == nil {
		return nil, errors.New("invalid value for required argument 'CreationData'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180930:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191101:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200501:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200630:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200930:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210801:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211201:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220302:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220702:Disk"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20230402:Disk"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Disk
	err := ctx.RegisterResource("azure-native:compute/v20230102:Disk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDisk gets an existing Disk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskState, opts ...pulumi.ResourceOption) (*Disk, error) {
	var resource Disk
	err := ctx.ReadResource("azure-native:compute/v20230102:Disk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Disk resources.
type diskState struct {
}

type DiskState struct {
}

func (DiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskState)(nil)).Elem()
}

type diskArgs struct {
	// Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is disabled by default. Does not apply to Ultra disks.
	BurstingEnabled *bool `pulumi:"burstingEnabled"`
	// Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent *float64 `pulumi:"completionPercent"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationData `pulumi:"creationData"`
	// Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode *string `pulumi:"dataAccessAuthMode"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `pulumi:"diskAccessId"`
	// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly *float64 `pulumi:"diskIOPSReadOnly"`
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadWrite *float64 `pulumi:"diskIOPSReadWrite"`
	// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly *float64 `pulumi:"diskMBpsReadOnly"`
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite *float64 `pulumi:"diskMBpsReadWrite"`
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskName *string `pulumi:"diskName"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption `pulumi:"encryption"`
	// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection `pulumi:"encryptionSettingsCollection"`
	// The extended location where the disk will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Resource location
	Location *string `pulumi:"location"`
	// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
	MaxShares *int `pulumi:"maxShares"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy *string `pulumi:"networkAccessPolicy"`
	// Setting this property to true improves reliability and performance of data disks that are frequently (more than 5 times a day) by detached from one virtual machine and attached to another. This property should not be set for disks that are not detached and attached frequently as it causes the disks to not align with the fault domain of the virtual machine.
	OptimizedForFrequentAttach *bool `pulumi:"optimizedForFrequentAttach"`
	// The Operating System type.
	OsType *OperatingSystemTypes `pulumi:"osType"`
	// Policy for controlling export on the disk.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Purchase plan information for the the image from which the OS disk was created. E.g. - {name: 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
	PurchasePlan *PurchasePlan `pulumi:"purchasePlan"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Contains the security related information for the resource.
	SecurityProfile *DiskSecurityProfile `pulumi:"securityProfile"`
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS, or PremiumV2_LRS.
	Sku *DiskSku `pulumi:"sku"`
	// List of supported capabilities for the image from which the OS disk was created.
	SupportedCapabilities *SupportedCapabilities `pulumi:"supportedCapabilities"`
	// Indicates the OS on a disk supports hibernation.
	SupportsHibernation *bool `pulumi:"supportsHibernation"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Performance tier of the disk (e.g, P4, S10) as described here: https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier *string `pulumi:"tier"`
	// The Logical zone list for Disk.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Disk resource.
type DiskArgs struct {
	// Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is disabled by default. Does not apply to Ultra disks.
	BurstingEnabled pulumi.BoolPtrInput
	// Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent pulumi.Float64PtrInput
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataInput
	// Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode pulumi.StringPtrInput
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId pulumi.StringPtrInput
	// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadOnly pulumi.Float64PtrInput
	// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
	DiskIOPSReadWrite pulumi.Float64PtrInput
	// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadOnly pulumi.Float64PtrInput
	// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
	DiskMBpsReadWrite pulumi.Float64PtrInput
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskName pulumi.StringPtrInput
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB pulumi.IntPtrInput
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption EncryptionPtrInput
	// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection EncryptionSettingsCollectionPtrInput
	// The extended location where the disk will be created. Extended location cannot be changed.
	ExtendedLocation ExtendedLocationPtrInput
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
	MaxShares pulumi.IntPtrInput
	// Policy for accessing the disk via network.
	NetworkAccessPolicy pulumi.StringPtrInput
	// Setting this property to true improves reliability and performance of data disks that are frequently (more than 5 times a day) by detached from one virtual machine and attached to another. This property should not be set for disks that are not detached and attached frequently as it causes the disks to not align with the fault domain of the virtual machine.
	OptimizedForFrequentAttach pulumi.BoolPtrInput
	// The Operating System type.
	OsType OperatingSystemTypesPtrInput
	// Policy for controlling export on the disk.
	PublicNetworkAccess pulumi.StringPtrInput
	// Purchase plan information for the the image from which the OS disk was created. E.g. - {name: 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
	PurchasePlan PurchasePlanPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Contains the security related information for the resource.
	SecurityProfile DiskSecurityProfilePtrInput
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS, or PremiumV2_LRS.
	Sku DiskSkuPtrInput
	// List of supported capabilities for the image from which the OS disk was created.
	SupportedCapabilities SupportedCapabilitiesPtrInput
	// Indicates the OS on a disk supports hibernation.
	SupportsHibernation pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Performance tier of the disk (e.g, P4, S10) as described here: https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier pulumi.StringPtrInput
	// The Logical zone list for Disk.
	Zones pulumi.StringArrayInput
}

func (DiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskArgs)(nil)).Elem()
}

type DiskInput interface {
	pulumi.Input

	ToDiskOutput() DiskOutput
	ToDiskOutputWithContext(ctx context.Context) DiskOutput
}

func (*Disk) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (i *Disk) ToDiskOutput() DiskOutput {
	return i.ToDiskOutputWithContext(context.Background())
}

func (i *Disk) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskOutput)
}

type DiskOutput struct{ *pulumi.OutputState }

func (DiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Disk)(nil)).Elem()
}

func (o DiskOutput) ToDiskOutput() DiskOutput {
	return o
}

func (o DiskOutput) ToDiskOutputWithContext(ctx context.Context) DiskOutput {
	return o
}

// Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is disabled by default. Does not apply to Ultra disks.
func (o DiskOutput) BurstingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolPtrOutput { return v.BurstingEnabled }).(pulumi.BoolPtrOutput)
}

// Latest time when bursting was last enabled on a disk.
func (o DiskOutput) BurstingEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.BurstingEnabledTime }).(pulumi.StringOutput)
}

// Percentage complete for the background copy when a resource is created via the CopyStart operation.
func (o DiskOutput) CompletionPercent() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.CompletionPercent }).(pulumi.Float64PtrOutput)
}

// Disk source information. CreationData information cannot be changed after the disk has been created.
func (o DiskOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v *Disk) CreationDataResponseOutput { return v.CreationData }).(CreationDataResponseOutput)
}

// Additional authentication requirements when exporting or uploading to a disk or snapshot.
func (o DiskOutput) DataAccessAuthMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.DataAccessAuthMode }).(pulumi.StringPtrOutput)
}

// ARM id of the DiskAccess resource for using private endpoints on disks.
func (o DiskOutput) DiskAccessId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.DiskAccessId }).(pulumi.StringPtrOutput)
}

// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
func (o DiskOutput) DiskIOPSReadOnly() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskIOPSReadOnly }).(pulumi.Float64PtrOutput)
}

// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
func (o DiskOutput) DiskIOPSReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskIOPSReadWrite }).(pulumi.Float64PtrOutput)
}

// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
func (o DiskOutput) DiskMBpsReadOnly() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskMBpsReadOnly }).(pulumi.Float64PtrOutput)
}

// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
func (o DiskOutput) DiskMBpsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.Float64PtrOutput { return v.DiskMBpsReadWrite }).(pulumi.Float64PtrOutput)
}

// The size of the disk in bytes. This field is read only.
func (o DiskOutput) DiskSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *Disk) pulumi.Float64Output { return v.DiskSizeBytes }).(pulumi.Float64Output)
}

// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
func (o DiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

// The state of the disk.
func (o DiskOutput) DiskState() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.DiskState }).(pulumi.StringOutput)
}

// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
func (o DiskOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v *Disk) EncryptionResponsePtrOutput { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
func (o DiskOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v *Disk) EncryptionSettingsCollectionResponsePtrOutput { return v.EncryptionSettingsCollection }).(EncryptionSettingsCollectionResponsePtrOutput)
}

// The extended location where the disk will be created. Extended location cannot be changed.
func (o DiskOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *Disk) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
func (o DiskOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Resource location
func (o DiskOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// A relative URI containing the ID of the VM that has the disk attached.
func (o DiskOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.ManagedBy }).(pulumi.StringOutput)
}

// List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
func (o DiskOutput) ManagedByExtended() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringArrayOutput { return v.ManagedByExtended }).(pulumi.StringArrayOutput)
}

// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
func (o DiskOutput) MaxShares() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.IntPtrOutput { return v.MaxShares }).(pulumi.IntPtrOutput)
}

// Resource name
func (o DiskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Policy for accessing the disk via network.
func (o DiskOutput) NetworkAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.NetworkAccessPolicy }).(pulumi.StringPtrOutput)
}

// Setting this property to true improves reliability and performance of data disks that are frequently (more than 5 times a day) by detached from one virtual machine and attached to another. This property should not be set for disks that are not detached and attached frequently as it causes the disks to not align with the fault domain of the virtual machine.
func (o DiskOutput) OptimizedForFrequentAttach() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolPtrOutput { return v.OptimizedForFrequentAttach }).(pulumi.BoolPtrOutput)
}

// The Operating System type.
func (o DiskOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.OsType }).(pulumi.StringPtrOutput)
}

// Properties of the disk for which update is pending.
func (o DiskOutput) PropertyUpdatesInProgress() PropertyUpdatesInProgressResponseOutput {
	return o.ApplyT(func(v *Disk) PropertyUpdatesInProgressResponseOutput { return v.PropertyUpdatesInProgress }).(PropertyUpdatesInProgressResponseOutput)
}

// The disk provisioning state.
func (o DiskOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Policy for controlling export on the disk.
func (o DiskOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Purchase plan information for the the image from which the OS disk was created. E.g. - {name: 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
func (o DiskOutput) PurchasePlan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v *Disk) PurchasePlanResponsePtrOutput { return v.PurchasePlan }).(PurchasePlanResponsePtrOutput)
}

// Contains the security related information for the resource.
func (o DiskOutput) SecurityProfile() DiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *Disk) DiskSecurityProfileResponsePtrOutput { return v.SecurityProfile }).(DiskSecurityProfileResponsePtrOutput)
}

// Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
func (o DiskOutput) ShareInfo() ShareInfoElementResponseArrayOutput {
	return o.ApplyT(func(v *Disk) ShareInfoElementResponseArrayOutput { return v.ShareInfo }).(ShareInfoElementResponseArrayOutput)
}

// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS, or PremiumV2_LRS.
func (o DiskOutput) Sku() DiskSkuResponsePtrOutput {
	return o.ApplyT(func(v *Disk) DiskSkuResponsePtrOutput { return v.Sku }).(DiskSkuResponsePtrOutput)
}

// List of supported capabilities for the image from which the OS disk was created.
func (o DiskOutput) SupportedCapabilities() SupportedCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v *Disk) SupportedCapabilitiesResponsePtrOutput { return v.SupportedCapabilities }).(SupportedCapabilitiesResponsePtrOutput)
}

// Indicates the OS on a disk supports hibernation.
func (o DiskOutput) SupportsHibernation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.BoolPtrOutput { return v.SupportsHibernation }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o DiskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Performance tier of the disk (e.g, P4, S10) as described here: https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
func (o DiskOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringPtrOutput { return v.Tier }).(pulumi.StringPtrOutput)
}

// The time when the disk was created.
func (o DiskOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o DiskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Unique Guid identifying the resource.
func (o DiskOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

// The Logical zone list for Disk.
func (o DiskOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Disk) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskOutput{})
}
