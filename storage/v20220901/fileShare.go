// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Properties of the file share, including Id, resource name, resource type, Etag.
type FileShare struct {
	pulumi.CustomResourceState

	// Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
	AccessTier pulumi.StringPtrOutput `pulumi:"accessTier"`
	// Indicates the last modification time for share access tier.
	AccessTierChangeTime pulumi.StringOutput `pulumi:"accessTierChangeTime"`
	// Indicates if there is a pending transition for access tier.
	AccessTierStatus pulumi.StringOutput `pulumi:"accessTierStatus"`
	// Indicates whether the share was deleted.
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// The deleted time if the share was deleted.
	DeletedTime pulumi.StringOutput `pulumi:"deletedTime"`
	// The authentication protocol that is used for the file share. Can only be specified when creating a share.
	EnabledProtocols pulumi.StringPtrOutput `pulumi:"enabledProtocols"`
	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Returns the date and time the share was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// Specifies whether the lease on a share is of infinite or fixed duration, only when the share is leased.
	LeaseDuration pulumi.StringOutput `pulumi:"leaseDuration"`
	// Lease state of the share.
	LeaseState pulumi.StringOutput `pulumi:"leaseState"`
	// The lease status of the share.
	LeaseStatus pulumi.StringOutput `pulumi:"leaseStatus"`
	// A name-value pair to associate with the share as metadata.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Remaining retention days for share that was soft deleted.
	RemainingRetentionDays pulumi.IntOutput `pulumi:"remainingRetentionDays"`
	// The property is for NFS share only. The default is NoRootSquash.
	RootSquash pulumi.StringPtrOutput `pulumi:"rootSquash"`
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400.
	ShareQuota pulumi.IntPtrOutput `pulumi:"shareQuota"`
	// The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.
	ShareUsageBytes pulumi.Float64Output `pulumi:"shareUsageBytes"`
	// List of stored access policies specified on the share.
	SignedIdentifiers SignedIdentifierResponseArrayOutput `pulumi:"signedIdentifiers"`
	// Creation time of share snapshot returned in the response of list shares with expand param "snapshots".
	SnapshotTime pulumi.StringOutput `pulumi:"snapshotTime"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The version of the share.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewFileShare registers a new resource with the given unique name, arguments, and options.
func NewFileShare(ctx *pulumi.Context,
	name string, args *FileShareArgs, opts ...pulumi.ResourceOption) (*FileShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230101:FileShare"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FileShare
	err := ctx.RegisterResource("azure-native:storage/v20220901:FileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileShare gets an existing FileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileShareState, opts ...pulumi.ResourceOption) (*FileShare, error) {
	var resource FileShare
	err := ctx.ReadResource("azure-native:storage/v20220901:FileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileShare resources.
type fileShareState struct {
}

type FileShareState struct {
}

func (FileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareState)(nil)).Elem()
}

type fileShareArgs struct {
	// Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
	AccessTier *string `pulumi:"accessTier"`
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The authentication protocol that is used for the file share. Can only be specified when creating a share.
	EnabledProtocols *string `pulumi:"enabledProtocols"`
	// Optional, used to expand the properties within share's properties. Valid values are: snapshots. Should be passed as a string with delimiter ','
	Expand *string `pulumi:"expand"`
	// A name-value pair to associate with the share as metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The property is for NFS share only. The default is NoRootSquash.
	RootSquash *string `pulumi:"rootSquash"`
	// The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ShareName *string `pulumi:"shareName"`
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400.
	ShareQuota *int `pulumi:"shareQuota"`
	// List of stored access policies specified on the share.
	SignedIdentifiers []SignedIdentifier `pulumi:"signedIdentifiers"`
}

// The set of arguments for constructing a FileShare resource.
type FileShareArgs struct {
	// Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
	AccessTier pulumi.StringPtrInput
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// The authentication protocol that is used for the file share. Can only be specified when creating a share.
	EnabledProtocols pulumi.StringPtrInput
	// Optional, used to expand the properties within share's properties. Valid values are: snapshots. Should be passed as a string with delimiter ','
	Expand pulumi.StringPtrInput
	// A name-value pair to associate with the share as metadata.
	Metadata pulumi.StringMapInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The property is for NFS share only. The default is NoRootSquash.
	RootSquash pulumi.StringPtrInput
	// The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ShareName pulumi.StringPtrInput
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400.
	ShareQuota pulumi.IntPtrInput
	// List of stored access policies specified on the share.
	SignedIdentifiers SignedIdentifierArrayInput
}

func (FileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareArgs)(nil)).Elem()
}

type FileShareInput interface {
	pulumi.Input

	ToFileShareOutput() FileShareOutput
	ToFileShareOutputWithContext(ctx context.Context) FileShareOutput
}

func (*FileShare) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (i *FileShare) ToFileShareOutput() FileShareOutput {
	return i.ToFileShareOutputWithContext(context.Background())
}

func (i *FileShare) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileShareOutput)
}

func (i *FileShare) ToOutput(ctx context.Context) pulumix.Output[*FileShare] {
	return pulumix.Output[*FileShare]{
		OutputState: i.ToFileShareOutputWithContext(ctx).OutputState,
	}
}

type FileShareOutput struct{ *pulumi.OutputState }

func (FileShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (o FileShareOutput) ToFileShareOutput() FileShareOutput {
	return o
}

func (o FileShareOutput) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return o
}

func (o FileShareOutput) ToOutput(ctx context.Context) pulumix.Output[*FileShare] {
	return pulumix.Output[*FileShare]{
		OutputState: o.OutputState,
	}
}

// Access tier for specific share. GpV2 account can choose between TransactionOptimized (default), Hot, and Cool. FileStorage account can choose Premium.
func (o FileShareOutput) AccessTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringPtrOutput { return v.AccessTier }).(pulumi.StringPtrOutput)
}

// Indicates the last modification time for share access tier.
func (o FileShareOutput) AccessTierChangeTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.AccessTierChangeTime }).(pulumi.StringOutput)
}

// Indicates if there is a pending transition for access tier.
func (o FileShareOutput) AccessTierStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.AccessTierStatus }).(pulumi.StringOutput)
}

// Indicates whether the share was deleted.
func (o FileShareOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v *FileShare) pulumi.BoolOutput { return v.Deleted }).(pulumi.BoolOutput)
}

// The deleted time if the share was deleted.
func (o FileShareOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.DeletedTime }).(pulumi.StringOutput)
}

// The authentication protocol that is used for the file share. Can only be specified when creating a share.
func (o FileShareOutput) EnabledProtocols() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringPtrOutput { return v.EnabledProtocols }).(pulumi.StringPtrOutput)
}

// Resource Etag.
func (o FileShareOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Returns the date and time the share was last modified.
func (o FileShareOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// Specifies whether the lease on a share is of infinite or fixed duration, only when the share is leased.
func (o FileShareOutput) LeaseDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.LeaseDuration }).(pulumi.StringOutput)
}

// Lease state of the share.
func (o FileShareOutput) LeaseState() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.LeaseState }).(pulumi.StringOutput)
}

// The lease status of the share.
func (o FileShareOutput) LeaseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.LeaseStatus }).(pulumi.StringOutput)
}

// A name-value pair to associate with the share as metadata.
func (o FileShareOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o FileShareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Remaining retention days for share that was soft deleted.
func (o FileShareOutput) RemainingRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *FileShare) pulumi.IntOutput { return v.RemainingRetentionDays }).(pulumi.IntOutput)
}

// The property is for NFS share only. The default is NoRootSquash.
func (o FileShareOutput) RootSquash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringPtrOutput { return v.RootSquash }).(pulumi.StringPtrOutput)
}

// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120). For Large File Shares, the maximum size is 102400.
func (o FileShareOutput) ShareQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileShare) pulumi.IntPtrOutput { return v.ShareQuota }).(pulumi.IntPtrOutput)
}

// The approximate size of the data stored on the share. Note that this value may not include all recently created or recently resized files.
func (o FileShareOutput) ShareUsageBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *FileShare) pulumi.Float64Output { return v.ShareUsageBytes }).(pulumi.Float64Output)
}

// List of stored access policies specified on the share.
func (o FileShareOutput) SignedIdentifiers() SignedIdentifierResponseArrayOutput {
	return o.ApplyT(func(v *FileShare) SignedIdentifierResponseArrayOutput { return v.SignedIdentifiers }).(SignedIdentifierResponseArrayOutput)
}

// Creation time of share snapshot returned in the response of list shares with expand param "snapshots".
func (o FileShareOutput) SnapshotTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.SnapshotTime }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FileShareOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The version of the share.
func (o FileShareOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FileShareOutput{})
}
