// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230402

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets information about a snapshot.
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSnapshotResult
	err := ctx.Invoke("azure-native:compute/v20230402:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80 characters.
	SnapshotName string `pulumi:"snapshotName"`
}

// Snapshot resource.
type LookupSnapshotResult struct {
	// Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent *float64 `pulumi:"completionPercent"`
	// Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
	CopyCompletionError *CopyCompletionErrorResponse `pulumi:"copyCompletionError"`
	// Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData CreationDataResponse `pulumi:"creationData"`
	// Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode *string `pulumi:"dataAccessAuthMode"`
	// ARM id of the DiskAccess resource for using private endpoints on disks.
	DiskAccessId *string `pulumi:"diskAccessId"`
	// The size of the disk in bytes. This field is read only.
	DiskSizeBytes float64 `pulumi:"diskSizeBytes"`
	// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `pulumi:"diskSizeGB"`
	// The state of the snapshot.
	DiskState string `pulumi:"diskState"`
	// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *EncryptionResponse `pulumi:"encryption"`
	// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	// The extended location where the snapshot will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *string `pulumi:"hyperVGeneration"`
	// Resource Id
	Id string `pulumi:"id"`
	// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
	Incremental *bool `pulumi:"incremental"`
	// Incremental snapshots for a disk share an incremental snapshot family id. The Get Page Range Diff API can only be called on incremental snapshots with the same family id.
	IncrementalSnapshotFamilyId string `pulumi:"incrementalSnapshotFamilyId"`
	// Resource location
	Location string `pulumi:"location"`
	// Unused. Always Null.
	ManagedBy string `pulumi:"managedBy"`
	// Resource name
	Name string `pulumi:"name"`
	// Policy for accessing the disk via network.
	NetworkAccessPolicy *string `pulumi:"networkAccessPolicy"`
	// The Operating System type.
	OsType *string `pulumi:"osType"`
	// The disk provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// Policy for controlling export on the disk.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Purchase plan information for the image from which the source disk for the snapshot was originally created.
	PurchasePlan *PurchasePlanResponse `pulumi:"purchasePlan"`
	// Contains the security related information for the resource.
	SecurityProfile *DiskSecurityProfileResponse `pulumi:"securityProfile"`
	// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
	Sku *SnapshotSkuResponse `pulumi:"sku"`
	// List of supported capabilities for the image from which the source disk from the snapshot was originally created.
	SupportedCapabilities *SupportedCapabilitiesResponse `pulumi:"supportedCapabilities"`
	// Indicates the OS on a snapshot supports hibernation.
	SupportsHibernation *bool `pulumi:"supportsHibernation"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The time when the snapshot was created.
	TimeCreated string `pulumi:"timeCreated"`
	// Resource type
	Type string `pulumi:"type"`
	// Unique Guid identifying the resource.
	UniqueId string `pulumi:"uniqueId"`
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResult, error) {
			args := v.(LookupSnapshotArgs)
			r, err := LookupSnapshot(ctx, &args, opts...)
			var s LookupSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotResultOutput)
}

type LookupSnapshotOutputArgs struct {
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the snapshot that is being created. The name can't be changed after the snapshot is created. Supported characters for the name are a-z, A-Z, 0-9, _ and -. The max name length is 80 characters.
	SnapshotName pulumi.StringInput `pulumi:"snapshotName"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}

// Snapshot resource.
type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSnapshotResult] {
	return pulumix.Output[LookupSnapshotResult]{
		OutputState: o.OutputState,
	}
}

// Percentage complete for the background copy when a resource is created via the CopyStart operation.
func (o LookupSnapshotResultOutput) CompletionPercent() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *float64 { return v.CompletionPercent }).(pulumi.Float64PtrOutput)
}

// Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
func (o LookupSnapshotResultOutput) CopyCompletionError() CopyCompletionErrorResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *CopyCompletionErrorResponse { return v.CopyCompletionError }).(CopyCompletionErrorResponsePtrOutput)
}

// Disk source information. CreationData information cannot be changed after the disk has been created.
func (o LookupSnapshotResultOutput) CreationData() CreationDataResponseOutput {
	return o.ApplyT(func(v LookupSnapshotResult) CreationDataResponse { return v.CreationData }).(CreationDataResponseOutput)
}

// Additional authentication requirements when exporting or uploading to a disk or snapshot.
func (o LookupSnapshotResultOutput) DataAccessAuthMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.DataAccessAuthMode }).(pulumi.StringPtrOutput)
}

// ARM id of the DiskAccess resource for using private endpoints on disks.
func (o LookupSnapshotResultOutput) DiskAccessId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.DiskAccessId }).(pulumi.StringPtrOutput)
}

// The size of the disk in bytes. This field is read only.
func (o LookupSnapshotResultOutput) DiskSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSnapshotResult) float64 { return v.DiskSizeBytes }).(pulumi.Float64Output)
}

// If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only allowed if the disk is not attached to a running VM, and can only increase the disk's size.
func (o LookupSnapshotResultOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

// The state of the snapshot.
func (o LookupSnapshotResultOutput) DiskState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.DiskState }).(pulumi.StringOutput)
}

// Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
func (o LookupSnapshotResultOutput) Encryption() EncryptionResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *EncryptionResponse { return v.Encryption }).(EncryptionResponsePtrOutput)
}

// Encryption settings collection used be Azure Disk Encryption, can contain multiple encryption settings per disk or snapshot.
func (o LookupSnapshotResultOutput) EncryptionSettingsCollection() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *EncryptionSettingsCollectionResponse {
		return v.EncryptionSettingsCollection
	}).(EncryptionSettingsCollectionResponsePtrOutput)
}

// The extended location where the snapshot will be created. Extended location cannot be changed.
func (o LookupSnapshotResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
func (o LookupSnapshotResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

// Resource Id
func (o LookupSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full snapshots and can be diffed.
func (o LookupSnapshotResultOutput) Incremental() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *bool { return v.Incremental }).(pulumi.BoolPtrOutput)
}

// Incremental snapshots for a disk share an incremental snapshot family id. The Get Page Range Diff API can only be called on incremental snapshots with the same family id.
func (o LookupSnapshotResultOutput) IncrementalSnapshotFamilyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.IncrementalSnapshotFamilyId }).(pulumi.StringOutput)
}

// Resource location
func (o LookupSnapshotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Location }).(pulumi.StringOutput)
}

// Unused. Always Null.
func (o LookupSnapshotResultOutput) ManagedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.ManagedBy }).(pulumi.StringOutput)
}

// Resource name
func (o LookupSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy for accessing the disk via network.
func (o LookupSnapshotResultOutput) NetworkAccessPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.NetworkAccessPolicy }).(pulumi.StringPtrOutput)
}

// The Operating System type.
func (o LookupSnapshotResultOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

// The disk provisioning state.
func (o LookupSnapshotResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Policy for controlling export on the disk.
func (o LookupSnapshotResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Purchase plan information for the image from which the source disk for the snapshot was originally created.
func (o LookupSnapshotResultOutput) PurchasePlan() PurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *PurchasePlanResponse { return v.PurchasePlan }).(PurchasePlanResponsePtrOutput)
}

// Contains the security related information for the resource.
func (o LookupSnapshotResultOutput) SecurityProfile() DiskSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *DiskSecurityProfileResponse { return v.SecurityProfile }).(DiskSecurityProfileResponsePtrOutput)
}

// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
func (o LookupSnapshotResultOutput) Sku() SnapshotSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *SnapshotSkuResponse { return v.Sku }).(SnapshotSkuResponsePtrOutput)
}

// List of supported capabilities for the image from which the source disk from the snapshot was originally created.
func (o LookupSnapshotResultOutput) SupportedCapabilities() SupportedCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *SupportedCapabilitiesResponse { return v.SupportedCapabilities }).(SupportedCapabilitiesResponsePtrOutput)
}

// Indicates the OS on a snapshot supports hibernation.
func (o LookupSnapshotResultOutput) SupportsHibernation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *bool { return v.SupportsHibernation }).(pulumi.BoolPtrOutput)
}

// Resource tags
func (o LookupSnapshotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSnapshotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The time when the snapshot was created.
func (o LookupSnapshotResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// Resource type
func (o LookupSnapshotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Type }).(pulumi.StringOutput)
}

// Unique Guid identifying the resource.
func (o LookupSnapshotResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
