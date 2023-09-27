// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230102

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about a disk.
func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDiskResult
	err := ctx.Invoke("azure-native:compute/v20230102:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskArgs struct {
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskName string `pulumi:"diskName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Disk resource.
type LookupDiskResult struct {
	// Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is disabled by default. Does not apply to Ultra disks.
	BurstingEnabled *bool `pulumi:"burstingEnabled"`
	// Latest time when bursting was last enabled on a disk.
	BurstingEnabledTime string `pulumi:"burstingEnabledTime"`
	// Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent *float64 `pulumi:"completionPercent"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponse `pulumi:"creationData"`
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
	// The size of the disk in bytes. This field is read only.
	DiskSizeBytes float64 `pulumi:"diskSizeBytes"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// The state of the disk.
	DiskState string `pulumi:"diskState"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *EncryptionResponse `pulumi:"encryption"`
	// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	// The extended location where the disk will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// A relative URI containing the ID of the VM that has the disk attached.
	ManagedBy string `pulumi:"managedBy"`
	// List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
	ManagedByExtended []string `pulumi:"managedByExtended"`
	// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
	MaxShares *int `pulumi:"maxShares"`
	// Resource name
	Name string `pulumi:"name"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy *string `pulumi:"networkAccessPolicy"`
	// Setting this property to true improves reliability and performance of data disks that are frequently (more than 5 times a day) by detached from one virtual machine and attached to another. This property should not be set for disks that are not detached and attached frequently as it causes the disks to not align with the fault domain of the virtual machine.
	OptimizedForFrequentAttach *bool `pulumi:"optimizedForFrequentAttach"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// Properties of the disk for which update is pending.
	PropertyUpdatesInProgress PropertyUpdatesInProgressResponse `pulumi:"propertyUpdatesInProgress"`
	// The disk provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Policy for controlling export on the disk.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Purchase plan information for the the image from which the OS disk was created. E.g. - {name: 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
	PurchasePlan *PurchasePlanResponse `pulumi:"purchasePlan"`
	// Contains the security related information for the resource.
	SecurityProfile *DiskSecurityProfileResponse `pulumi:"securityProfile"`
	// Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
	ShareInfo []ShareInfoElementResponse `pulumi:"shareInfo"`
	// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS, or PremiumV2_LRS.
	Sku *DiskSkuResponse `pulumi:"sku"`
	// List of supported capabilities for the image from which the OS disk was created.
	SupportedCapabilities *SupportedCapabilitiesResponse `pulumi:"supportedCapabilities"`
	// Indicates the OS on a disk supports hibernation.
	SupportsHibernation *bool `pulumi:"supportsHibernation"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Performance tier of the disk (e.g, P4, S10) as described here: https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
	Tier *string `pulumi:"tier"`
	// The time when the disk was created.
	TimeCreated string `pulumi:"timeCreated"`
	// Resource type
	Type string `pulumi:"type"`
	// Unique Guid identifying the resource.
	UniqueId string `pulumi:"uniqueId"`
	// The Logical zone list for Disk.
	Zones []string `pulumi:"zones"`
}

func LookupDiskOutput(ctx *pulumi.Context, args LookupDiskOutputArgs, opts ...pulumi.InvokeOption) LookupDiskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiskResult, error) {
			args := v.(LookupDiskArgs)
			r, err := LookupDisk(ctx, &args, opts...)
			var s LookupDiskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiskResultOutput)
}

type LookupDiskOutputArgs struct {
	// The name of the managed disk that is being created. The name can't be changed after the disk is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The maximum name length is 80 characters.
	DiskName pulumi.StringInput `pulumi:"diskName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDiskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskArgs)(nil)).Elem()
}

// Disk resource.
type LookupDiskResultOutput struct{ *pulumi.OutputState }

func (LookupDiskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiskResult)(nil)).Elem()
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutput() LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) ToLookupDiskResultOutputWithContext(ctx context.Context) LookupDiskResultOutput {
	return o
}

func (o LookupDiskResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDiskResult] {
	return pulumix.Output[LookupDiskResult]{
		OutputState: o.OutputState,
	}
}

// Set to true to enable bursting beyond the provisioned performance target of the disk. Bursting is disabled by default. Does not apply to Ultra disks.
func (o LookupDiskResultOutput) BurstingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *bool { return v.BurstingEnabled }).(pulumi.BoolPtrOutput)
}

// Latest time when bursting was last enabled on a disk.
func (o LookupDiskResultOutput) BurstingEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.BurstingEnabledTime }).(pulumi.StringOutput)
}

// Percentage complete for the background copy when a resource is created via the CopyStart operation.
func (o LookupDiskResultOutput) CompletionPercent() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.CompletionPercent }).(pulumi.Float64PtrOutput)
}

// Disk source information. CreationData information cannot be changed after the disk has been created.
func (o LookupDiskResultOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) CreationDataResponse { return v.CreationData }).(CreationDataResponseOutput)
}

// Additional authentication requirements when exporting or uploading to a disk or snapshot.
func (o LookupDiskResultOutput) DataAccessAuthMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.DataAccessAuthMode }).(pulumi.StringPtrOutput)
}

// ARM id of the DiskAccess resource for using private endpoints on disks.
func (o LookupDiskResultOutput) DiskAccessId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.DiskAccessId }).(pulumi.StringPtrOutput)
}

// The total number of IOPS that will be allowed across all VMs mounting the shared disk as ReadOnly. One operation can transfer between 4k and 256k bytes.
func (o LookupDiskResultOutput) DiskIOPSReadOnly() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.DiskIOPSReadOnly }).(pulumi.Float64PtrOutput)
}

// The number of IOPS allowed for this disk; only settable for UltraSSD disks. One operation can transfer between 4k and 256k bytes.
func (o LookupDiskResultOutput) DiskIOPSReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.DiskIOPSReadWrite }).(pulumi.Float64PtrOutput)
}

// The total throughput (MBps) that will be allowed across all VMs mounting the shared disk as ReadOnly. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
func (o LookupDiskResultOutput) DiskMBpsReadOnly() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.DiskMBpsReadOnly }).(pulumi.Float64PtrOutput)
}

// The bandwidth allowed for this disk; only settable for UltraSSD disks. MBps means millions of bytes per second - MB here uses the ISO notation, of powers of 10.
func (o LookupDiskResultOutput) DiskMBpsReadWrite() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *float64 { return v.DiskMBpsReadWrite }).(pulumi.Float64PtrOutput)
}

// The size of the disk in bytes. This field is read only.
func (o LookupDiskResultOutput) DiskSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupDiskResult) float64 { return v.DiskSizeBytes }).(pulumi.Float64Output)
}

// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
func (o LookupDiskResultOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

// The state of the disk.
func (o LookupDiskResultOutput) DiskState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.DiskState }).(pulumi.StringOutput)
}

// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
func (o LookupDiskResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// Encryption settings collection used for Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
func (o LookupDiskResultOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *EncryptionSettingsCollectionResponse { return v.EncryptionSettingsCollection }).(EncryptionSettingsCollectionResponsePtrOutput)
}

// The extended location where the disk will be created. Extended location cannot be changed.
func (o LookupDiskResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
func (o LookupDiskResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupDiskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupDiskResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Location }).(pulumi.StringOutput)
}

// A relative URI containing the ID of the VM that has the disk attached.
func (o LookupDiskResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

// List of relative URIs containing the IDs of the VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
func (o LookupDiskResultOutput) ManagedByExtended() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.ManagedByExtended }).(pulumi.StringArrayOutput)
}

// The maximum number of VMs that can attach to the disk at the same time. Value greater than one indicates a disk that can be mounted on multiple VMs at the same time.
func (o LookupDiskResultOutput) MaxShares() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *int { return v.MaxShares }).(pulumi.IntPtrOutput)
}

// Resource name
func (o LookupDiskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy for accessing the disk via network.
func (o LookupDiskResultOutput) NetworkAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.NetworkAccessPolicy }).(pulumi.StringPtrOutput)
}

// Setting this property to true improves reliability and performance of data disks that are frequently (more than 5 times a day) by detached from one virtual machine and attached to another. This property should not be set for disks that are not detached and attached frequently as it causes the disks to not align with the fault domain of the virtual machine.
func (o LookupDiskResultOutput) OptimizedForFrequentAttach() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *bool { return v.OptimizedForFrequentAttach }).(pulumi.BoolPtrOutput)
}

// The Operating System type.
func (o LookupDiskResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

// Properties of the disk for which update is pending.
func (o LookupDiskResultOutput) PropertyUpdatesInProgress() PropertyUpdatesInProgressResponseOutput {
	return o.ApplyT(func(v LookupDiskResult) PropertyUpdatesInProgressResponse { return v.PropertyUpdatesInProgress }).(PropertyUpdatesInProgressResponseOutput)
}

// The disk provisioning state.
func (o LookupDiskResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Policy for controlling export on the disk.
func (o LookupDiskResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Purchase plan information for the the image from which the OS disk was created. E.g. - {name: 2019-Datacenter, publisher: MicrosoftWindowsServer, product: WindowsServer}
func (o LookupDiskResultOutput) PurchasePlan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *PurchasePlanResponse { return v.PurchasePlan }).(PurchasePlanResponsePtrOutput)
}

// Contains the security related information for the resource.
func (o LookupDiskResultOutput) SecurityProfile() DiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *DiskSecurityProfileResponse { return v.SecurityProfile }).(DiskSecurityProfileResponsePtrOutput)
}

// Details of the list of all VMs that have the disk attached. maxShares should be set to a value greater than one for disks to allow attaching them to multiple VMs.
func (o LookupDiskResultOutput) ShareInfo() ShareInfoElementResponseArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []ShareInfoElementResponse { return v.ShareInfo }).(ShareInfoElementResponseArrayOutput)
}

// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, UltraSSD_LRS, Premium_ZRS, StandardSSD_ZRS, or PremiumV2_LRS.
func (o LookupDiskResultOutput) Sku() DiskSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *DiskSkuResponse { return v.Sku }).(DiskSkuResponsePtrOutput)
}

// List of supported capabilities for the image from which the OS disk was created.
func (o LookupDiskResultOutput) SupportedCapabilities() SupportedCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *SupportedCapabilitiesResponse { return v.SupportedCapabilities }).(SupportedCapabilitiesResponsePtrOutput)
}

// Indicates the OS on a disk supports hibernation.
func (o LookupDiskResultOutput) SupportsHibernation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *bool { return v.SupportsHibernation }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o LookupDiskResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDiskResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Performance tier of the disk (e.g, P4, S10) as described here: https://azure.microsoft.com/en-us/pricing/details/managed-disks/. Does not apply to Ultra disks.
func (o LookupDiskResultOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiskResult) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

// The time when the disk was created.
func (o LookupDiskResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o LookupDiskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.Type }).(pulumi.StringOutput)
}

// Unique Guid identifying the resource.
func (o LookupDiskResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiskResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

// The Logical zone list for Disk.
func (o LookupDiskResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDiskResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiskResultOutput{})
}
